package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// var Collection *mongo.Collection

const connectionString = "mongodb+srv://Naman:Naman@12345@cluster0.bezaxfc.mongodb.net/?retryWrites=true&w=majority"
const DbName = "NetFlix"
const ColName = "watchlist"

func ConnectDB() {

	//client option
	clientOption := options.Client().ApplyURI(connectionString)
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

}
