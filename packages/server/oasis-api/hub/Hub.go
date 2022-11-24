package hub

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type MessageFeed struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ContactId string `json:"contactId"`
}

type Time struct {
	Seconds string `json:"seconds"`
	Nanos   string `json:"nanos"`
}

type MessageItem struct {
	Username  string `json:"username"`
	FeedId    string `json:"feedId"`
	Id        string `json:"id"`
	Direction string `json:"direction"`
	Message   string `json:"message"`
	Time      Time   `json:"time"`
	Channel   string `json:"channel"`
}

type FeedHub struct {
	Clients       map[*websocket.Conn]bool
	FeedBroadcast chan MessageFeed
	Sync          *sync.Cond
}

type MessageHub struct {
	Clients          map[string]map[*websocket.Conn]bool
	MessageBroadcast chan MessageItem
	Sync             *sync.Cond
}

func NewFeedHub() *FeedHub {
	return &FeedHub{
		Clients:       make(map[*websocket.Conn]bool),
		FeedBroadcast: make(chan MessageFeed),
		Sync:          sync.NewCond(new(sync.Mutex)),
	}
}

func NewMessageHub() *MessageHub {
	return &MessageHub{
		Clients:          make(map[string]map[*websocket.Conn]bool),
		MessageBroadcast: make(chan MessageItem),
		Sync:             sync.NewCond(new(sync.Mutex)),
	}
}

func (h *MessageHub) RunMessageHub() {
	for {
		select {
		case message := <-h.MessageBroadcast:
			log.Printf("Message Hub: got message for feed id %s", message.FeedId)
			if message.Id == "quit" {
				log.Printf("Message Hub: Got the kill command, shutting down")
				h.MessageBroadcast <- MessageItem{}
				return
			}
			if conns := h.Clients[message.FeedId]; conns != nil {
				for conn := range conns {
					log.Printf("Sending message to Webscoket")
					if err := conn.WriteJSON(message); !errors.Is(err, nil) {
						log.Printf("error occurred: %v", err)
					}
				}
			}
		}
	}
}

func (h *FeedHub) RunFeedHub() {
	for {
		select {
		case feed := <-h.FeedBroadcast:
			if feed.ContactId == "quit" {
				log.Printf("Feed Hub: Got the kill command, shutting down")
				h.FeedBroadcast <- MessageFeed{}
				return
			}
			for client := range h.Clients {
				if client != nil {
					if err := client.WriteJSON(feed); !errors.Is(err, nil) {
						log.Printf("error occurred: %v", err)
					}
				}
			}
		}
	}
}
