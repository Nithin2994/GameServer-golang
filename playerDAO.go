package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func updateLevel(playerName string, level int) string {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	filter := bson.M{"playerName": playerName}
	update := bson.D{
		{"$set", bson.M{"playerLevel": level}},
	}
	_, err = client.Database("server").Collection("PlayerInfo").UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))

	if err != nil {
		log.Fatal(err)
		return failed
	}
	return success
}

func getPlayerLevel(playerName string) int {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	var player playerInfo

	filter := bson.M{"playerName": playerName}

	err = client.Database("server").Collection("P	layerInfo").FindOne(context.Background(), filter).Decode(&player)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	return player.PlayerLevel
}
