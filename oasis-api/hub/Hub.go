package hub

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
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
}

type MessageHub struct {
	Clients          map[string]*websocket.Conn
	MessageBroadcast chan MessageItem
}

func NewFeedHub() *FeedHub {
	return &FeedHub{
		Clients:       make(map[*websocket.Conn]bool),
		FeedBroadcast: make(chan MessageFeed),
	}
}

func NewMessageHub() *MessageHub {
	return &MessageHub{
		Clients:          make(map[string]*websocket.Conn),
		MessageBroadcast: make(chan MessageItem),
	}
}

func (h *MessageHub) RunMessageHub() {
	for {
		select {
		case message := <-h.MessageBroadcast:
			if message.Id == "quit" {
				log.Printf("Message Hub: Got the kill command, shutting down")
				h.MessageBroadcast <- MessageItem{}
				return
			}
			if conn := h.Clients[message.FeedId]; conn != nil {
				if err := conn.WriteJSON(message); !errors.Is(err, nil) {
					log.Printf("error occurred: %v", err)
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
