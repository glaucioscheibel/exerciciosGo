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
	cursor, _ := collection.Find(
		ctx,
		bson.M{
			"genres": bson.M{
				"$all": []string{"Drama", "Comedy"},
			},
		},
		options.Find().SetProjection(
			bson.M{
				"title":  1,
				"genres": 1,
				"_id":    0,
			},
		),
	)
	for cursor.Next(ctx) {
		var doc bson.M
		cursor.Decode(&doc)
		jsonBytes, _ := json.MarshalIndent(doc, "", "  ")
		fmt.Printf("Movie: %s\n\n", string(jsonBytes))
	}
	cursor.Close(ctx)
	client.Disconnect(ctx)
}
