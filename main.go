package main

import "mongo_demo/db"

//

func main() {
	//var err error
	//var collection *mongo.Collection
	//collection, err = db.ConnectDB()
	//if err != nil {
	//	fmt.Printf("connect to db failed, err:%v\n", err)
	//}

	//notice := db.ReadFromNte()
	//err = db.InsertDataToNtcs(collection, notice)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//}

	//err = db.DisconnectDB()
	//if err != nil {
	//	fmt.Printf("connect to db failed, err:%v\n", err)
	//}
	db.ReadFromNte()
}
