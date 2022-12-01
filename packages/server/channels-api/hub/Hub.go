package hub

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type WebChatMessageItem struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Kill     bool   `json:"-"`
}

type WebChatMessageHub struct {
	Clients          map[string]map[*websocket.Conn]bool
	MessageBroadcast chan WebChatMessageItem
	Sync             *sync.Cond
}

func NewWebChatMessageHub() *WebChatMessageHub {
	return &WebChatMessageHub{
		Clients:          make(map[string]map[*websocket.Conn]bool),
		MessageBroadcast: make(chan WebChatMessageItem),
		Sync:             sync.NewCond(new(sync.Mutex)),
	}
}

func (h *WebChatMessageHub) RunWebChatMessageHub(pingInterval time.Duration) {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()
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
		case <-ticker.C:
			log.Printf("Sending pings for WebChat hub")
			for username := range h.Clients {
				for conn := range h.Clients[username] {
					conn.SetWriteDeadline(time.Now().Add(pingInterval / 2))
					if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						log.Printf("Ping Failed on websocket")
					}
				}
			}

		}
	}
}
