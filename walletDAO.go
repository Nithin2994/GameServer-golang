package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	balances := make(map[string]int64)
	balances["xp"] = 300
	balances["hc"] = 30
	updateBalances("Nithin", balances)
	updatedBalances := getBalances("Nithin")
	fmt.Println(updatedBalances)
	balance := getBalance("Nithin", "xp")
	fmt.Println(balance)
	status := transaction("Nithin", "xp", 200, "CR")
	fmt.Println("transaction : " + status)
	status = transaction("Nithin", "xp", 800, "DB")
	fmt.Println("transaction : " + status)
}

func updateBalances(playerName string, balances map[string]int64) string {

	filter := bson.D{{"playerName", playerName}}
	update := bson.D{
		{"$set", bson.M{"playerName": playerName, "balances": balances}},
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return failed
	}

	collection := client.Database("server").Collection("PlayerWallet")
	_, err = collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
		return failed
	}

	fmt.Println("Connection to MongoDB closed.")
	return success
}

func getBalances(playerName string) map[string]int64 {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	filter := bson.D{{"playerName", playerName}}
	var playerWallet playerWallet
	collection := client.Database("server").Collection("PlayerWallet")
	err = collection.FindOne(context.Background(), filter).Decode(&playerWallet)
	if err != nil {
		log.Fatal(err)
	}
	return playerWallet.Balances
}

func getBalance(playerName string, currency string) int64 {
	return getBalances(playerName)[currency]
}

func transaction(playerName string, currency string, amount int64, transactionType string) string {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	filter := bson.D{{"playerName", playerName}}
	collection := client.Database("server").Collection("PlayerWallet")

	balance := getBalance(playerName, currency)

	if transactionType == "CR" {
		update := bson.D{
			{"$inc", bson.M{"balances." + currency: amount}},
		}
		_, err = collection.UpdateOne(context.Background(), filter, update)
	} else if transactionType == "DB" {
		if balance < amount {
			return "Insufficient Funds"
		}

		update := bson.D{
			{"$inc", bson.M{"balances." + currency: -amount}},
		}
		_, err = collection.UpdateOne(context.Background(), filter, update)
	} else {
		return "invalid transaction type"
	}
	if err != nil {
		log.Fatal(err)
	}
	return "success"
}
