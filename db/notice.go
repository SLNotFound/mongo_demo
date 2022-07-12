package db

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mongo_demo/model"
	"os"
	"strconv"
)

// 读取通知文件内容

func ReadFromNte() *model.Notice {
	f, err := os.Open("./test.nte")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	var dataList []string
	for scanner.Scan() {
		dataList = append(dataList, scanner.Text())
	}
	var params = make([]string, 5)
	params = dataList[1:6]
	propsMap := make(map[string]string)
	propsMap["ack"] = dataList[7]
	propsMap["ackparam"] = dataList[9]
	propsMap["subject"] = dataList[23]
	propsMap["msgFlag"] = dataList[17]
	propsMap["msgtype"] = dataList[19]
	propsMap["datapath"] = dataList[11]
	method, _ := strconv.Atoi(dataList[0])

	notice := &model.Notice{
		Method:  method,
		PcRead:  0,
		MobRead: 0,
		SendId:  dataList[2],
		RecvId:  "",
		MsgId:   dataList[1],
		Params:  params,
		Props:   propsMap,
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return notice
}

func InsertDataToNtcs(collection *mongo.Collection, notice *model.Notice) error {
	result, err := collection.InsertOne(context.TODO(), &notice)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)
	return nil
}
