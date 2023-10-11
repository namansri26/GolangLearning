package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naman/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Route = gin.Default()
var Client *mongo.Client

func RouteURl() {
	Route.POST("/register", InsertOneMovie)

}

func main() {
	// database.ConnectDB()

	//client option
	clientOption := options.Client().ApplyURI("mongodb+srv://Naman:Naman%4012345@cluster0.bezaxfc.mongodb.net/?retryWrites=true&w=majority")
	Client, err := mongo.Connect(context.Background(), clientOption)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Check if the connection is established
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		return
	}
	fmt.Println("MongoDb connection success")
	fmt.Println("Hello form main Mongo Db APis")
	RouteURl()

	err = Route.Run(":8080")

	if err != nil {
		log.Fatal(err)
		return
	}

}

func init() {
	// database.ConnectDB()

}
func InsertOneMovie(ctx *gin.Context) {
	var request model.NetFlix
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Fatal(err)
		return
	}
	collection := Client.Database("NetFlix").Collection("watchlist")
	fmt.Println("COLLECTION : ", collection)
	inserted, err := collection.InsertOne(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)

	fmt.Println("Inserted one Id in db: ", inserted.InsertedID)

}
