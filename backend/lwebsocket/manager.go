package lwebsocket

import (
	"fmt"
)

type WebsocketManager struct {
	Clients    map[*Client]bool
	Broadcast  chan *Message
	Register   chan *Client
	Unregister chan *Client
}

func NewWebsocketManager() *WebsocketManager {
	wsm := WebsocketManager{
		Broadcast:  make(chan *Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
	go wsm.start()
	return &wsm
}

type Message struct {
	Itemid  string
	Content string
}

func (manager *WebsocketManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			fmt.Println("new ws")
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				fmt.Println("unregister ws")
			}

		case message := <-manager.Broadcast:
			for client := range manager.Clients {
				if client.Itemid == message.Itemid {
					client.Send <- *message
				}
			}
		}
	}
}
