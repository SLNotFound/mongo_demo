package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mongo_demo/model"
)

func InsertDataToMsgs(collection *mongo.Collection, msgList []*model.Msg) {
	for _, msg := range msgList {
		result, err := collection.InsertOne(context.TODO(), &msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.InsertedID)
	}
}
