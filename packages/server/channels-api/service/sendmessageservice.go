package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	mimemail "github.com/emersion/go-message/mail"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	proto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes/chatHub"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	oryClient "github.com/ory/client-go"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"strings"
	"time"
)

type sendMessageService struct {
	proto.UnimplementedMessageEventServiceServer
	conf        *c.Config
	mh          *chatHub.Hub
	df          util.DialFactory
	oauthConfig *oauth2.Config
}

func getIdentityIdMetadataForGRPC(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata")
	}

	kh := md.Get("X-Openline-IDENTITY-ID")
	if kh != nil && len(kh) == 1 {
		return kh[0], nil
	}
	return "", errors.New("no IdentityId header")
}

func (s sendMessageService) SendMessageEvent(c context.Context, msg *msProto.InputMessage) (*msProto.MessageId, error) {
	username, err := service.GetUsernameMetadataForGRPC(c)
	if err != nil {
		log.Printf("Missing username header")
		return nil, err
	}

	identityId, err := getIdentityIdMetadataForGRPC(c)
	if err != nil {
		log.Printf("Missing Identity Id header")
		return nil, err
	}

	log.Printf("Got a and ID Hader of %s", identityId)
	conn, err := s.df.GetMessageStoreCon()
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return nil, err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)

	switch msg.Type {
	case msProto.MessageType_EMAIL:
		newBody, mailErr := s.sendMail(identityId, msg)
		if mailErr != nil {
			return nil, mailErr
		}
		bytes, err := json.Marshal(newBody)
		if err != nil {
			return nil, err
		}
		bodyStr := string(bytes)
		msg.Content = &bodyStr
	case msProto.MessageType_WEB_CHAT:
		webChatErr := s.sendWebChat(*username, msg)
		if webChatErr != nil {
			return nil, webChatErr
		}
	default:
		err := fmt.Errorf("unknown channel: %s", msg.Type)
		return nil, err
	}

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, s.conf.Service.MessageStoreApiKey)
	ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, *username)

	newMsg, err := client.SaveMessage(ctx, msg)
	if err != nil {
		log.Printf("Unable to connect to retrieve message!")
		return nil, err
	}
	return newMsg, nil
}

func (s sendMessageService) sendWebChat(username string, msg *msProto.InputMessage) error {

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, s.conf.Service.MessageStoreApiKey)
	ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, username)
	conn, err := s.df.GetMessageStoreCon()
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)

	participants, err := client.GetParticipants(ctx, &msProto.FeedId{Id: *msg.ConversationId})
	if err != nil {
		log.Printf("Unable to get participants from message store!")
		return err
	}

	for _, participant := range participants.Participants {
		if participant == username {
			continue
		}
		// Send a message to the hub
		messageItem := chatHub.MessageItem{
			Username: participant,
			Message:  *msg.Content,
		}

		s.mh.Broadcast <- messageItem
		log.Printf("successfully sent new message for %s", participant)
	}

	return nil
}

func (s sendMessageService) getMailAuthToken(identityId string) (*oauth2.Token, error) {
	configuration := oryClient.NewConfiguration()
	configuration.Servers = []oryClient.ServerConfiguration{
		{
			URL: s.conf.GMail.OryServerUrl,
		},
	}
	ory := oryClient.NewAPIClient(configuration)
	ctx := context.Background()
	ctx = context.WithValue(ctx, oryClient.ContextAccessToken, s.conf.GMail.OryApiKey)
	identity, _, err := ory.IdentityApi.GetIdentity(ctx, identityId).IncludeCredential([]string{"oidc"}).Execute()
	if err != nil {
		log.Printf("Unable to get gmail auth token for %s, (%s)", identityId, err.Error())
		return nil, err
	}
	credentials := identity.GetCredentials()["oidc"]
	log.Printf("Got credentials of %v", credentials)

	providers, ok := credentials.GetConfig()["providers"].([]interface{})
	log.Printf("Got providers of %T", providers[0])

	if !ok {
		log.Printf("unable to get provider list %s", identityId)
		return nil, err
	}

	provider, ok := providers[0].(map[string]interface{})
	if !ok {
		log.Printf("unable to get provider list %s", identityId)
		return nil, err
	}
	token, ok := provider["initial_access_token"].(string)

	if !ok {
		log.Printf("unable to get access token %s", identityId)
		return nil, err
	}
	tok := &oauth2.Token{AccessToken: token, TokenType: "Bearer"}

	refreshToken, ok := provider["initial_refresh_token"].(string)

	if !ok {
		log.Printf("unable to get refresh token`` %s", identityId)
	} else {
		log.Printf("Setting refresh token to %s", refreshToken)
		tok.RefreshToken = refreshToken
	}
	tok.Expiry = time.Now().Add(time.Hour * -1)
	return tok, nil
}

func (s sendMessageService) sendMail(identityId string, msg *msProto.InputMessage) (*routes.EmailContent, error) {
	tok, err := s.getMailAuthToken(identityId)

	log.Printf("Got Auth Token of %v", tok)
	client := s.oauthConfig.Client(context.Background(), tok)

	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))

	jsonMail := &routes.EmailContent{}
	err = json.Unmarshal([]byte(*msg.Content), jsonMail)
	if err != nil {
		log.Printf("Unable to parse email content for %s", msg.InitiatorIdentifier)
		return nil, err
	}
	fromAddress := []*mimemail.Address{{"", jsonMail.From}}
	toAddress := []*mimemail.Address{}
	for _, to := range jsonMail.To {
		toAddress = append(toAddress, &mimemail.Address{"", to})
	}

	var b bytes.Buffer
	user := "me"

	// Create our mail header
	var h mimemail.Header
	h.SetDate(time.Now())
	h.SetAddressList("From", fromAddress)
	h.SetAddressList("To", toAddress)
	h.SetMessageID(jsonMail.MessageId)
	h.SetSubject(jsonMail.Subject)

	// Create a new mail writer
	mw, err := mimemail.CreateWriter(&b, h)
	if err != nil {
		log.Fatal(err)
	}

	// Create a text part
	tw, err := mw.CreateInline()
	if err != nil {
		log.Fatal(err)
	}
	var th mimemail.InlineHeader
	th.Set("Content-Type", "text/html")
	w, err := tw.CreatePart(th)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, jsonMail.Html)
	w.Close()
	tw.Close()

	mw.Close()

	raw := base64.StdEncoding.EncodeToString(b.Bytes())
	msgToSend := &gmail.Message{
		Raw: raw,
	}
	result, err := srv.Users.Messages.Send(user, msgToSend).Do()
	if err != nil {
		log.Printf("Unable to send email: %v", err)
		return nil, err
	}

	generatedMessage, err := srv.Users.Messages.Get("me", result.Id).Do()
	if err != nil {
		log.Printf("Unable to get email: %v", err)
		return nil, err
	}
	for _, header := range generatedMessage.Payload.Headers {
		log.Printf("Comparing %s to %s", header.Name, "Message-ID")
		if strings.EqualFold(header.Name, "Message-ID") {
			jsonMail.MessageId = header.Value
			break
		}
	}
	log.Printf("Email successfully sent id %v", jsonMail.MessageId)
	return jsonMail, nil
}

//
//	smtpClient, err := s.df.GetSMTPClientCon()
//	if err != nil {
//		log.Printf("Unable to connect to mail server! %v", err)
//		return err
//	}
//
//	// Create email
//	email := mail.NewMSG()
//	email.SetFrom(s.conf.Mail.SMTPFromUser)
//	email.AddTo(msg.GetUsername())
//	email.SetSubject("Hello") //TODO
//
//	email.SetBody(mail.TextPlain, msg.GetMessage())
//
//	err = email.Send(smtpClient)
//	if err != nil {
//		log.Printf("Unable to send to mail server!")
//		return err
//	}
//	log.Printf("Email successfully sent to %s", msg.GetUsername())
//	return nil
//}

func NewSendMessageService(c *c.Config, df util.DialFactory, oauthConfig *oauth2.Config, mh *chatHub.Hub) *sendMessageService {
	ms := new(sendMessageService)
	ms.conf = c
	ms.mh = mh
	ms.df = df
	ms.oauthConfig = oauthConfig
	return ms
}
