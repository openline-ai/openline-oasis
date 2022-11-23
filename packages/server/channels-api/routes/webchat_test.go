package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"net/http/httptest"
	"openline-ai/channels-api/config"
	"openline-ai/channels-api/test_utils"
	"openline-ai/oasis-api/proto"
	"testing"
	"time"
)

var webchatRouter *gin.Engine

const webchatApiKey = "89ab1b51-0d03-4ce2-942b-ce0445188563"

func init() {
	dft := test_utils.MakeDialFactoryTest()
	webchatRouter = gin.Default()
	route := webchatRouter.Group("/")

	conf := &config.Config{}
	conf.WebChat.ApiKey = webchatApiKey
	AddWebChatRoutes(conf, dft, route)
}

func TestSendWebchat(t *testing.T) {
	id1 := int64(7)
	gabyId := int64(2)
	sentMessageEvent := false

	mpr := &WebchatMessage{
		Username: "gabi@example.org",
		Message:  "Hello Gabi",
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

	test_utils.SetChannelApiCallbacks(&test_utils.MockChannelApiCallbacks{NewMessageEvent: func(ctx context.Context, id *proto.OasisMessageId) (*proto.OasisEmpty, error) {
		if !assert.Equal(t, id.MessageId, id1) {
			return nil, status.Error(400, "Unexpected message id!")
		}
		sentMessageEvent = true
		return &proto.OasisEmpty{}, nil
	}})

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(mpr)
	req, _ := http.NewRequest("POST", "/webchat/", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("WebChatApiKey", webchatApiKey)
	log.Printf("Going to post body %s", reqBody)
	webchatRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)
	if !assert.Equal(t, w.Code, 200) {
		return
	}
	assert.True(t, sentMessageEvent, "NewMessageEvent not called!")
}

func TestSendWebchatInvalidApiKey(t *testing.T) {
	id1 := int64(7)
	gabyId := int64(2)
	sentMessageEvent := false
	sentSaveMessage := false

	mpr := &WebchatMessage{
		Username: "gabi@example.org",
		Message:  "Hello Gabi",
	}
	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{SaveMessage: func(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
		log.Printf("Inside SaveMessage")
		sentSaveMessage = true
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

	test_utils.SetChannelApiCallbacks(&test_utils.MockChannelApiCallbacks{NewMessageEvent: func(ctx context.Context, id *proto.OasisMessageId) (*proto.OasisEmpty, error) {
		if !assert.Equal(t, id.MessageId, id1) {
			return nil, status.Error(400, "Unexpected message id!")
		}
		sentMessageEvent = true
		return &proto.OasisEmpty{}, nil
	}})

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(mpr)
	req, _ := http.NewRequest("POST", "/webchat/", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	log.Printf("Going to post body %s", reqBody)
	webchatRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)
	if !assert.Equal(t, w.Code, 403) {
		return
	}
	assert.False(t, sentMessageEvent, "NewMessageEvent called when it shouldn't!")
	assert.False(t, sentSaveMessage, "NewMessageEvent called when it shouldn't!")

}

func TestKeyValidator(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/webchat/?email=gabi@example.org", nil)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("WebChatApiKey", webchatApiKey)
	webchatRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)
	if !assert.Equal(t, w.Code, 200) {
		return
	}
	var resp LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("TestKeyValidator: %v\n", err.Error())
		return
	}
	assert.Equal(t, "gabi@example.org", resp.UserName)
}

func TestKeyValidatorInvalidKey(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/webchat/?email=gabi@example.org", nil)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	webchatRouter.ServeHTTP(w, req)
	log.Printf("Got Body %s", w.Body)
	if !assert.Equal(t, w.Code, 403) {
		return
	}

}
