package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}

func (h *ItemsHandler) AddItem(c *gin.Context) {
	itempostreq := &ItemPostrequest{}
	err := c.BindJSON(itempostreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var email, _ = c.Get("my_mail")
	user, err := h.Userstore.GetByMail(email.(string))
	item := types.Item{
		Name:        itempostreq.Name,
		Description: itempostreq.Description,
		UserID:      user.ID,
		ImageUrl:    itempostreq.ImageUrl,
	}
	h.Itemstore.Insert(&item)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, item.ToResponse())
}
func (h *ItemsHandler) GetAllItems(c *gin.Context) {
	items, err := h.Itemstore.GetAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, types.ItemsToResponse(*items))
}

func (h *ItemsHandler) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := h.Itemstore.GetById(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, item.ToResponse())
}

type BidReq struct {
	Price float32 `json:"price,omitempty"`
}

func (h *ItemsHandler) AddBid(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	bidreq := &BidReq{}
	err := c.BindJSON(bidreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var email, _ = c.Get("my_mail")
	user, err := h.Userstore.GetByMail(email.(string))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	bid := &types.Bid{UserID: user.ID, Price: bidreq.Price, ItemID: id}
	err = h.Itemstore.InsertBid(bid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	bidresp, _ := h.Itemstore.GetBidById(bid.ID)
	bidjson, err := json.Marshal(bidresp.ToResponse())
	fmt.Println(string(bidjson))
	if err != nil {
		return
	}

	h.WebsocketManager.Broadcast <- &lwebsocket.Message{
		Itemid:  strconv.Itoa(id),
		Content: string(bidjson),
	}
	c.Status(200)
}
