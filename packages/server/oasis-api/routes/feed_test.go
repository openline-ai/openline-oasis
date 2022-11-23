package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"net/http/httptest"
	chanProto "openline-ai/channels-api/ent/proto"
	"openline-ai/oasis-api/config"
	"openline-ai/oasis-api/test_utils"
	"testing"
	"time"
)

var feedRouter *gin.Engine

func init() {
	dft := test_utils.MakeDialFactoryTest()
	feedRouter = gin.Default()
	route := feedRouter.Group("/")

	addFeedRoutes(route, &config.Config{}, dft)

}

func TestGetFeeds(t *testing.T) {
	resp := msProto.FeedList{}
	id1 := int64(0)
	id2 := int64(1)

	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{GetFeeds: func(ctx context.Context, empty *msProto.Empty) (*msProto.FeedList, error) {
		fl := &msProto.FeedList{Contact: make([]*msProto.Contact, 2)}

		fl.Contact[0] = &msProto.Contact{ContactId: "12345678", FirstName: "Torrey", LastName: "Searle", Id: &id1}
		fl.Contact[1] = &msProto.Contact{ContactId: "77775553", FirstName: "Gabriel", LastName: "Gontariu", Id: &id2}
		return fl, nil
	}})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/feed", nil)
	feedRouter.ServeHTTP(w, req)
	if !assert.Equal(t, w.Code, 200) {
		return
	}

	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("TestGetFeeds: %v\n", err)
		return
	}
	if resp.GetContact() == nil {
		t.Errorf("missing list!")
	}
	assert.Equal(t, resp.GetContact()[0].FirstName, "Torrey")
	assert.Equal(t, resp.GetContact()[0].LastName, "Searle")
	assert.Equal(t, *resp.GetContact()[0].Id, id1)
	assert.Equal(t, resp.GetContact()[0].ContactId, "12345678")

	assert.Equal(t, resp.GetContact()[1].FirstName, "Gabriel")
	assert.Equal(t, resp.GetContact()[1].LastName, "Gontariu")
	assert.Equal(t, *resp.GetContact()[1].Id, id2)
	assert.Equal(t, resp.GetContact()[1].ContactId, "77775553")

	log.Printf("Got a response of %v\n", resp)

}

func TestGetFeed(t *testing.T) {
	resp := &msProto.Contact{}
	id1 := int64(0)
	id2 := int64(1)
	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{GetFeed: func(ctx context.Context, contact *msProto.Contact) (*msProto.Contact, error) {
		log.Printf("Entering GetFeed: %v", contact)
		if !assert.NotNil(t, contact.Id, "we expected the id from the path to be passed") {
			return nil, status.Error(400, "Contact ID should not be null")
		}

		if *contact.Id == int64(0) {
			return &msProto.Contact{ContactId: "12345678", FirstName: "Torrey", LastName: "Searle", Id: &id1}, nil
		} else if *contact.Id == int64(1) {
			return &msProto.Contact{ContactId: "77775553", FirstName: "Gabriel", LastName: "Gontariu", Id: &id2}, nil
		}
		return nil, status.Errorf(codes.Unknown, "Error Finding Feed")
	}})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/feed/0", nil)
	feedRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)
	if !assert.Equal(t, w.Code, 200) {
		return
	}

	err := json.Unmarshal(w.Body.Bytes(), &resp)

	if err != nil {
		t.Errorf("TestGetFeed: %v\n", err)
		return
	}
	assert.Equal(t, resp.FirstName, "Torrey")
	assert.Equal(t, resp.LastName, "Searle")
	assert.Equal(t, *resp.Id, id1)
	assert.Equal(t, resp.ContactId, "12345678")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/feed/1", nil)
	feedRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)
	if !assert.Equal(t, w.Code, 200) {
		return
	}

	err = json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, resp.FirstName, "Gabriel")
	assert.Equal(t, resp.LastName, "Gontariu")
	assert.Equal(t, *resp.Id, id2)
	assert.Equal(t, resp.ContactId, "77775553")

}

func TestGetMessages(t *testing.T) {
	gabyId := int64(2)
	id1 := int64(7)
	id2 := int64(15)
	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{GetMessages: func(ctx context.Context, contact *msProto.PagedContact) (*msProto.MessageList, error) {
		if !assert.NotNil(t, *contact.Contact.Id) {
			return nil, status.Error(400, "Contact ID should not be null")
		}
		if !assert.Equal(t, *contact.Contact.Id, gabyId) {
			return nil, status.Errorf(400, "Contact ID should should be 2 but got %d", *contact.Contact.Id)
		}

		ml := &msProto.MessageList{Message: make([]*msProto.Message, 2)}
		mi := &msProto.Message{
			Type:      msProto.MessageType_MESSAGE, //MESSAGE
			Message:   "Hello Gabi",
			Direction: msProto.MessageDirection_INBOUND, // INBOUND
			Channel:   msProto.MessageChannel_MAIL,      // Mail
			Username:  "gabi@example.org",
			Id:        &id1,
			Time:      timestamppb.New(time.Now()),
			Contact:   &msProto.Contact{ContactId: "77775553"},
		}
		ml.Message[0] = mi
		mi = &msProto.Message{
			Type:      msProto.MessageType_MESSAGE, //MESSAGE
			Message:   "Hey there, how are you?",
			Direction: msProto.MessageDirection_OUTBOUND, // OUTBOUND
			Channel:   msProto.MessageChannel_MAIL,       // Mail
			Username:  "gabi@example.org",
			Id:        &id2,
			Time:      timestamppb.New(time.Now()),
			Contact:   &msProto.Contact{ContactId: "77775553"},
		}
		ml.Message[1] = mi
		return ml, nil

	}})

	var resp []*msProto.Message
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/feed/"+fmt.Sprint(gabyId)+"/item", nil)
	feedRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)

	if !assert.Equal(t, w.Code, 200) {
		return
	}

	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("TestGetMessages: %v\n", err)
		return
	}
	if !assert.Equal(t, len(resp), 2) {
		return
	}

	log.Printf("Message 1:\n%v", *resp[0])
	assert.Equal(t, resp[0].Type, msProto.MessageType_MESSAGE)
	assert.Equal(t, resp[0].GetMessage(), "Hello Gabi")
	assert.Equal(t, resp[0].Direction, msProto.MessageDirection_INBOUND)
	assert.Equal(t, resp[0].Channel, msProto.MessageChannel_MAIL)
	assert.Equal(t, resp[0].GetUsername(), "gabi@example.org")
	assert.Equal(t, *resp[0].Id, id1)

	log.Printf("Message 1:\n%v", *resp[1])

	assert.Equal(t, resp[1].Type, msProto.MessageType_MESSAGE)
	assert.Equal(t, resp[1].GetMessage(), "Hey there, how are you?")
	assert.Equal(t, resp[1].Direction, msProto.MessageDirection_OUTBOUND)
	assert.Equal(t, resp[1].Channel, msProto.MessageChannel_MAIL)
	assert.Equal(t, resp[1].GetUsername(), "gabi@example.org")
	assert.Equal(t, *resp[1].Id, id2)

}

func TestSaveMessages(t *testing.T) {
	id1 := int64(7)
	gabyId := int64(2)

	mi := &msProto.Message{
		Type:      msProto.MessageType_MESSAGE, //MESSAGE
		Message:   "Hello Gabi",
		Direction: msProto.MessageDirection_INBOUND, // INBOUND
		Channel:   msProto.MessageChannel_MAIL,      // Mail
		Username:  "gabi@example.org",
		Id:        &id1,
		Time:      nil,
		Contact:   &msProto.Contact{ContactId: "77775553"},
	}
	fpr := &FeedPostRequest{
		Username:  "gabi@example.org",
		Message:   "Hello Gabi",
		Channel:   "MAIL",
		Source:    "WEB",
		Direction: "INBOUND",
	}
	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{SaveMessage: func(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
		log.Printf("Inside SaveMessage")
		var tm *time.Time = nil
		if message.GetTime() != nil {
			var timeref = message.GetTime().AsTime()
			tm = &timeref
		}

		if tm == nil {
			var timeRef = time.Now()
			tm = &timeRef
		}

		message.Time = timestamppb.New(*tm)
		message.Id = &id1

		if message.Contact == nil {
			if !assert.Equal(t, "gabi@example.org", message.Username) {
				return nil, status.Error(400, "Unexpected username")
			}
		} else {
			if !assert.Equal(t, gabyId, *message.Contact.Id) {
				return nil, status.Error(400, "Unexpected contact ID")
			}
		}

		message.Contact = &msProto.Contact{
			ContactId: "77775553",
			FirstName: "Gabriel",
			LastName:  "Gontariu",
			Id:        &gabyId,
		}
		return message, nil
	}})

	test_utils.SetChannelApiCallbacks(&test_utils.MockChannelApi{SendMessageEvent: func(ctx context.Context, id *chanProto.MessageId) (*chanProto.EventEmpty, error) {
		if !assert.Equal(t, id.MessageId, id1) {
			return nil, status.Error(400, "Unexpected message id!")
		}
		return &chanProto.EventEmpty{}, nil
	}})

	var resp msProto.Message
	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(fpr)
	req, _ := http.NewRequest("POST", "/feed/"+fmt.Sprint(gabyId)+"/item", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	log.Printf("Going to post body %s", reqBody)
	feedRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)
	if !assert.Equal(t, w.Code, 200) {
		return
	}

	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("TestGetMessages: %v\n", err)
		return
	}
	assert.Equal(t, mi.Message, resp.Message)
}
