package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yasngleer/bidex/store"
	"github.com/yasngleer/bidex/types"
)

type UserHandler struct {
	store        store.UserStore
	sessionstore store.SessionStore
}

func NewUserHandler(userstore store.UserStore, sessionstore store.SessionStore) *UserHandler {
	return &UserHandler{store: userstore, sessionstore: sessionstore}
}

type UserRegisterRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *UserHandler) UserRegister(c *gin.Context) {
	uregreq := &UserRegisterRequest{}
	err := c.BindJSON(uregreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, _ := types.NewUser(uregreq.Email, uregreq.Password)
	err = h.store.Insert(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	token := uuid.New().String()
	h.sessionstore.Insert(token, uregreq.Email)
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.Status(200)
}
func (u *UserHandler) UserLogin(c *gin.Context) {
	uregreq := &UserRegisterRequest{}
	err := c.BindJSON(uregreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	usr, _ := u.store.GetByMail(uregreq.Email)
	if usr.ValidatePassword(uregreq.Password) {
		token := uuid.New().String()

		u.sessionstore.Insert(token, uregreq.Email)
		c.SetCookie("token", token, 3600, "/", "localhost", false, true)

		c.Status(200)

		return
	}
	c.Status(401)

}
func (u *UserHandler) MeHandler(c *gin.Context) {
	var data, exists = c.Get("my_mail")

	if exists {
		c.String(200, data.(string))
	} else {
		c.Status(400)
		return

	}

}
func (u *UserHandler) Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	u.sessionstore.Delete(cookie.Value)

}
