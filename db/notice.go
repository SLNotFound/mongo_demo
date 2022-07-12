package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo_demo/model"
)

func InsertDataToNtcs(collection *mongo.Collection, noticeList []*model.Notice) error {
	for _, notice := range noticeList {
		result, err := collection.InsertOne(context.TODO(), &notice)
		if err != nil {
			return err
		}
		fmt.Println(result.InsertedID)
	}
	//result, err := collection.InsertOne(context.TODO(), &noticeList)
	//if err != nil {
	//	return err
	//}

	return nil
}
