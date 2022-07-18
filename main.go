package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo_demo/db"
	"mongo_demo/util/msg"
	"mongo_demo/util/notice"
)

func main() {
	var err error
	var c1 *mongo.Collection
	var c2 *mongo.Collection
	err = db.ConnectDB()
	if err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
	}

	nte := notice.ReadFromNte()
	c1 = db.Client.Database("notice").Collection("ntcs")
	db.InsertDataToNtcs(c1, nte)

	msgList, date := msg.ReadDataFromMsg()
	c2 = db.Client.Database("messages").Collection("msg-" + date)
	db.InsertDataToMsgs(c2, msgList)

	defer db.DisconnectDB()
}
