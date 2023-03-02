package FeedHub

import (
	"encoding/json"
	"log"
)

type ReloadFeed struct {
}

// FeedHub Hub maintains the set of active clients and broadcasts messages to the
// clients.
type FeedHub struct {
	// Registered clients.
	Clients map[*FeedClient]bool

	// Inbound messages from the clients.
	Broadcast chan ReloadFeed

	// Register requests from the clients.
	register chan *FeedClient

	// Unregister requests from clients.
	unregister chan *FeedClient

	Quit chan bool
}

func NewFeedHub() *FeedHub {
	return &FeedHub{
		Broadcast:  make(chan ReloadFeed),
		register:   make(chan *FeedClient),
		unregister: make(chan *FeedClient),
		Clients:    make(map[*FeedClient]bool),
		Quit:       make(chan bool),
	}
}

func (h *FeedHub) Run() {
	for {
		select {
		case quit := <-h.Quit:
			if quit {
				log.Printf("Kill request received, shutting down")
				return
			}
		case client := <-h.register:
			h.Clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				byteMsg, err := json.Marshal(message)
				if err != nil {
					log.Printf("FeedHub: Unable to marchal feed %s", err)
					continue
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
