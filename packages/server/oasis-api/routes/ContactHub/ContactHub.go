package ContactHub

import (
	"encoding/json"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	"log"
)

type Time struct {
	Seconds string `json:"seconds"`
	Nanos   string `json:"nanos"`
}

type ContactEvent struct {
	ContactId string           `json:"feedId"`
	Message   *msProto.Message `json:"message"`
}

// ContactHub Hub maintains the set of active clients and broadcasts messages to the
// clients.
type ContactHub struct {
	// Registered clients.
	Clients map[*ContactClient]bool

	// Inbound messages from the clients.
	Broadcast chan ContactEvent

	// Register requests from the clients.
	Register chan *ContactClient

	// Unregister requests from clients.
	unregister chan *ContactClient

	Quit chan bool
}

func NewContactHub() *ContactHub {
	return &ContactHub{
		Broadcast:  make(chan ContactEvent),
		Register:   make(chan *ContactClient),
		unregister: make(chan *ContactClient),
		Clients:    make(map[*ContactClient]bool),
		Quit:       make(chan bool),
	}
}

func (h *ContactHub) Run() {
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
			log.Printf("Unable to marchal message for feed: %s", message.ContactId)
			for client := range h.Clients {
				log.Printf("Unable to marchal message for feed: %s", message.ContactId)
				if client.feedId == message.ContactId {
					byteMsg, err := json.Marshal(message.Message)
					if err != nil {
						log.Printf("Unable to marchal message for feed: %s", message.ContactId)
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
