package model

type Notice struct {
	Method     int      `bson:"method"`
	CreateTime int64    `bson:"createTime"`
	PcRead     int      `bson:"pcRead"`
	MobRead    int      `bson:"mobRead"`
	SendId     string   `bson:"sendId"`
	RecvId     string   `bson:"recvId"`
	MsgId      string   `bson:"msgId"`
	Params     []string `bson:"params"`

	Props map[string]string `bson:"props"`
}
