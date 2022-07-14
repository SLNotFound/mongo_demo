package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func ConnectDB() (err error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	//collection = client.Database("notice").Collection("ntcs")
	return nil
}

func DisconnectDB() {
	_ = Client.Disconnect(context.TODO())
}
