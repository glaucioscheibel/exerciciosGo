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
		bson.D{
			bson.E{
				Key: "$and",
				Value: bson.A{
					bson.D{bson.E{Key: "year", Value: bson.D{bson.E{Key: "$gte", Value: 1980}}}},
					bson.D{bson.E{Key: "year", Value: bson.D{bson.E{Key: "$lt", Value: 1990}}}},
				},
			},
		},
		options.Find().SetSort(
			bson.D{
				bson.E{Key: "imdb.rating", Value: -1},
			},
		),
	)
	for cursor.Next(ctx) {
		var doc bson.M
		cursor.Decode(&doc)
		jsonBytes, _ := json.MarshalIndent(doc, "", "  ")
		fmt.Printf("Achei: %s\n\n", string(jsonBytes))
	}
	cursor.Close(ctx)
	client.Disconnect(ctx)
}
