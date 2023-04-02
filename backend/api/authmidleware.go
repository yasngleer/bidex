package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasngleer/bidex/store"
)

func AuthMiddleware(sessionstore store.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		email, err := sessionstore.Get(cookie.Value)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("my_mail", email)

	}
}
