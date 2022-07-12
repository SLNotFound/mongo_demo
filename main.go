package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo_demo/db"
	"mongo_demo/util/notice"
)

func main() {
	var err error
	var collection *mongo.Collection
	collection, err = db.ConnectDB()
	if err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
	}

	nte := notice.ReadFromNte()

	err = db.InsertDataToNtcs(collection, nte)
	if err != nil {
		fmt.Println("err: ", err)
	}

	err = db.DisconnectDB()
	if err != nil {
		fmt.Printf("disconnect to db failed, err:%v\n", err)
	}

}
