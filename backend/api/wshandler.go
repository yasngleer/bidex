package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yasngleer/bidex/lwebsocket"
)

type WsHandler struct {
	Wsmanager *lwebsocket.WebsocketManager
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *WsHandler) Ws(c *gin.Context) {
	id := c.Param("id")
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)

	client := lwebsocket.Client{
		Itemid: id,
		Socket: conn,
		Send:   make(chan lwebsocket.Message),
	}

	h.Wsmanager.Register <- &client
	client.Write()
}
