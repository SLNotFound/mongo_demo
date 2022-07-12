package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mongo_demo/model"
)

func InsertDataToNtcs(collection *mongo.Collection, noticeList []*model.Notice) {
	for _, notice := range noticeList {
		result, err := collection.InsertOne(context.TODO(), &notice)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.InsertedID)
	}
}
