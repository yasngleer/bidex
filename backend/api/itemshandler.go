package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasngleer/bidex/lwebsocket"
	"github.com/yasngleer/bidex/store"
	"github.com/yasngleer/bidex/types"
)

type ItemsHandler struct {
	Itemstore        store.ItemStore
	Userstore        store.UserStore
	WebsocketManager *lwebsocket.WebsocketManager
}

func NewItemsHandler(itemstore store.ItemStore, userstore store.UserStore, WebsocketManager *lwebsocket.WebsocketManager) *ItemsHandler {
	return &ItemsHandler{
		Itemstore:        itemstore,
		Userstore:        userstore,
		WebsocketManager: WebsocketManager,
	}
}

type ItemPostrequest struct {
	Name        string `json:"name,omitempty" bson:"name"`
	Description string `json:"description,omitempty" bson:"description"`
	ImageUrl    string `json:"image_url,omitempty" bson:"image_url"`
}

func (h *ItemsHandler) AddItem(c *gin.Context) {
	itempostreq := &ItemPostrequest{}
	err := c.BindJSON(itempostreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var email, _ = c.Get("my_mail")
	user, err := h.Userstore.GetByMail(c, email.(string))
	item := &types.Items{
		Name:        itempostreq.Name,
		Description: itempostreq.Description,
		UserID:      user.ID,
		ImageUrl:    itempostreq.ImageUrl,
	}
	h.Itemstore.Insert(c, item)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, item)
}
func (h *ItemsHandler) GetAllItems(c *gin.Context) {
	items, err := h.Itemstore.GetAll(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, items)
}

func (h *ItemsHandler) GetItem(c *gin.Context) {
	id := c.Param("id")
	items, err := h.Itemstore.GetById(c, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, items)
}

type BidReq struct {
	Price float32 `json:"price,omitempty" bson:"price"`
}

func (h *ItemsHandler) AddBid(c *gin.Context) {
	id := c.Param("id")
	bidreq := &BidReq{}
	err := c.BindJSON(bidreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var email, _ = c.Get("my_mail")
	user, err := h.Userstore.GetByMail(c, email.(string))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	bid := &types.Bid{UserID: user.ID, Price: bidreq.Price}
	err = h.Itemstore.InsertBid(c, id, bid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	bid.User = []types.User{{Email: user.Email, ID: user.ID}}
	bidjson, err := json.Marshal(bid)
	if err != nil {
		return
	}

	h.WebsocketManager.Broadcast <- &lwebsocket.Message{
		Itemid:  id,
		Content: string(bidjson),
	}
	c.JSON(200, "")
}
