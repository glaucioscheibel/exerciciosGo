package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	database := client.Database("mflix")
	collection := database.Collection("movies")

	cursor, _ := collection.Find(ctx, bson.M{"title": "The Room"})

	for cursor.Next(ctx) {
		var doc bson.M
		cursor.Decode(&doc)
		fmt.Println(doc)
	}
	cursor.Close(ctx)
	client.Disconnect(ctx)
}
