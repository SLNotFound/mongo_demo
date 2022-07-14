package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

func ConnectDB() (collection *mongo.Collection, err error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("notice").Collection("ntcs")
	return collection, nil
}

func DisconnectDB() {
	_ = client.Disconnect(context.TODO())
}
