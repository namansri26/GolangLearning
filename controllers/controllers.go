package controller

// connect with MongoDb

//Mongodb Helpers

//Insert one record

// func InsertOneMovie(ctx *gin.Context) {
// 	var request model.NetFlix
// 	err := ctx.ShouldBindJSON(&request)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	collection := .Client.Database(database.DbName).Collection(database.ColName)
// 	fmt.Println("COLLECTION : ", collection)
// 	inserted, err := collection.InsertOne(context.Background(), request)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	ctx.Writer.WriteHeader(http.StatusOK)

// 	fmt.Println("Inserted one Id in db: ", inserted.InsertedID)

// }

// func UpdateOneMove(movieid string) {
// 	id, err := primitive.ObjectIDFromHex(movieid)
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": bson.M{"isWatched": true}}
// 	result, err := collection.UpdateOne(context.background(), filter, update)

// 	if err != nil {
// 		Log.Fatal(err)

// 	}

// 	fmt.PrintLn("modified count :", result.ModifiedCount)

// }
