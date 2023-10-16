package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naman/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

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
