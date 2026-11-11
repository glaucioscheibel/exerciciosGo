package main

import (
	"context"
	"encoding/json"
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
	cursor, _ := collection.Find(ctx, bson.M{"runtime": bson.M{"$lt": 15}})
	for cursor.Next(ctx) {
		var doc bson.M
		cursor.Decode(&doc)
		jsonBytes, _ := json.MarshalIndent(doc, "", "  ")
		fmt.Printf("Movie: %s\n\n", string(jsonBytes))
	}
	cursor.Close(ctx)
	client.Disconnect(ctx)
}
