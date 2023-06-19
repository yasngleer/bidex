package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yasngleer/bidex/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yasngleer/bidex/lwebsocket"
	"github.com/yasngleer/bidex/store"
)

func main() {
	dbsql, err := gorm.Open(sqlite.Open("test_file.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:8080"}

	r.Use(cors.New(config))

	v1 := r.Group("/api")

	redissessiontore := store.NewRedisSessionStore("redis://localhost:6379")
	gormuserstore := store.NewGormUserStore(dbsql)
	gormitemstore := store.NewGormItemStore(dbsql)
	websocketmanager := lwebsocket.NewWebsocketManager()
	wshandler := &api.WsHandler{Wsmanager: websocketmanager}
	userhandler := api.NewUserHandler(gormuserstore, redissessiontore)
	itemhandler := api.NewItemsHandler(gormitemstore, gormuserstore, websocketmanager)
	v1.POST("/users", userhandler.UserRegister)
	v1.POST("/users/login", userhandler.UserLogin)
	v1.GET("/ws/:id", wshandler.Ws)

	//with auth
	v1.Use(api.AuthMiddleware(redissessiontore))
	v1.GET("/user/me", userhandler.MeHandler)
	v1.GET("/user/logout", userhandler.Logout)

	v1.POST("/items", itemhandler.AddItem)
	v1.GET("/items", itemhandler.GetAllItems)
	v1.GET("/items/:id", itemhandler.GetItem)

	v1.POST("/items/:id/bid", itemhandler.AddBid)

	r.Run(":8000")
}
