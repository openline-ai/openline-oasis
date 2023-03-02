package MessageHub

import (
	"encoding/json"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	"log"
)

type Time struct {
	Seconds string `json:"seconds"`
	Nanos   string `json:"nanos"`
}

type MessageEvent struct {
	FeedId  string           `json:"feedId"`
	Message *msProto.Message `json:"message"`
}

// MessageHub Hub maintains the set of active clients and broadcasts messages to the
// clients.
type MessageHub struct {
	// Registered clients.
	Clients map[*MessageClient]bool

	// Inbound messages from the clients.
	Broadcast chan MessageEvent

	// Register requests from the clients.
	Register chan *MessageClient

	// Unregister requests from clients.
	unregister chan *MessageClient

	Quit chan bool
}

func NewMessageHub() *MessageHub {
	return &MessageHub{
		Broadcast:  make(chan MessageEvent),
		Register:   make(chan *MessageClient),
		unregister: make(chan *MessageClient),
		Clients:    make(map[*MessageClient]bool),
		Quit:       make(chan bool),
	}
}

func (h *MessageHub) Run() {
	for {
		select {
		case quit := <-h.Quit:
			if quit {
				log.Printf("Kill request received, shutting down")
				return
			}
		case client := <-h.Register:
			log.Printf("Registered!")
			h.Clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			log.Printf("Unable to marchal message for feed: %s", message.FeedId)
			for client := range h.Clients {
				log.Printf("Unable to marchal message for feed: %s", message.FeedId)
				if client.feedId == message.FeedId {
					byteMsg, err := json.Marshal(message.Message)
					if err != nil {
						log.Printf("Unable to marchal message for feed: %s", message.FeedId)
						return
					}
					select {
					case client.send <- byteMsg:
					default:
						close(client.send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}
