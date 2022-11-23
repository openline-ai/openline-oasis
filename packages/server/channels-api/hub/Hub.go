package hub

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
)

type WebChatMessageItem struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Kill     bool   `json:"-"`
}

type WebChatMessageHub struct {
	Clients          map[string]map[*websocket.Conn]bool
	MessageBroadcast chan WebChatMessageItem
}

func NewWebChatMessageHub() *WebChatMessageHub {
	return &WebChatMessageHub{
		Clients:          make(map[string]map[*websocket.Conn]bool),
		MessageBroadcast: make(chan WebChatMessageItem),
	}
}

func (h *WebChatMessageHub) RunWebChatMessageHub() {
	for {
		select {
		case message := <-h.MessageBroadcast:
			if message.Kill {
				log.Printf("Kill request received, shutting down")
				h.MessageBroadcast <- WebChatMessageItem{}
				return
			}
			if conns := h.Clients[message.Username]; conns != nil {
				for conn := range conns {
					if err := conn.WriteJSON(message); !errors.Is(err, nil) {
						log.Printf("error occurred: %v", err.Error())
					}
				}
			}
		}
	}
}
