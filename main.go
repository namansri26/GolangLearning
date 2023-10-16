package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	// for creating the models
	controller "github.com/naman/mongoapi/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Route = gin.Default()
var Client *mongo.Client

func main() {
	// database.ConnectDB()

	//client option
	gin.SetMode(gin.ReleaseMode)
	Route := gin.New()
	var err error
	clientOption := options.Client().ApplyURI("mongodb+srv://Naman:Naman%4012345@cluster0.bezaxfc.mongodb.net/?retryWrites=true&w=majority")
	Client, err = mongo.Connect(context.Background(), clientOption)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Check if the connection is established or not
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		return
	}
	fmt.Println("MongoDb connection success")
	fmt.Println("Hello from main Mongo Db APis")

	Route.POST("/register", controller.InsertOneMovie)

	err = Route.Run(":8080")

	if err != nil {
		log.Fatal(err)
		return
	}

}
