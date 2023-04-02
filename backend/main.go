package main

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yasngleer/bidex/api"

	"github.com/yasngleer/bidex/lwebsocket"
	"github.com/yasngleer/bidex/store"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:8080"}

	r.Use(cors.New(config))

	v1 := r.Group("/api")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		panic(err)
	}
	db := client.Database("bidex")
	redissessiontore := store.NewRedisSessionStore("redis://redis")
	mongouserstore := store.NewMongoUserStore(db)
	mongoitemstore := store.NewMongoItemStore(db)
	websocketmanager := lwebsocket.NewWebsocketManager()
	wshandler := &api.WsHandler{Wsmanager: websocketmanager}
	userhandler := api.NewUserHandler(mongouserstore, redissessiontore)
	itemhandler := api.NewItemsHandler(mongoitemstore, mongouserstore, websocketmanager)
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
