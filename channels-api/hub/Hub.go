package hub

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
)

type WebChatMessageItem struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type WebChatMessageHub struct {
	Clients          map[string]*websocket.Conn
	MessageBroadcast chan WebChatMessageItem
}

func NewWebChatMessageHub() *WebChatMessageHub {
	return &WebChatMessageHub{
		Clients:          make(map[string]*websocket.Conn),
		MessageBroadcast: make(chan WebChatMessageItem),
	}
}

func (h *WebChatMessageHub) RunWebChatMessageHub() {
	for {
		select {
		case message := <-h.MessageBroadcast:
			if conn := h.Clients[message.Username]; conn != nil {
				if err := conn.WriteJSON(message); !errors.Is(err, nil) {
					log.Printf("error occurred: %v", err)
				}
			}
		}
	}
}
