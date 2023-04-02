package lwebsocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Itemid string
	//Connected socket
	Socket *websocket.Conn
	//Message
	Send chan Message
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message := <-c.Send:
			_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(message.Content))
		}
	}
}
