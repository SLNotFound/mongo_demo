package model

type Msg struct {
	MsgId        string `bson:"msg_id"`
	Subject      string `bson:"subject"`
	SendId       string `bson:"send_id"`
	SendName     string `bson:"send_name"`
	ReceiveId    string `bson:"receive_id"`
	ReceiveName  string `bson:"receive_name"`
	ReceiveUsers string `bson:"receive_users"`
	SendIp       string `bson:"send_ip"`
	SendMac      string `bson:"send_mac"`
	DataPath     string `bson:"data_path"`
	ContentType  string `bson:"content_type"`
	ExtData      string `bson:"ext_data"`
	SourceId     string `bson:"source_id"`
	Receiver     string `bson:"receiver"`
	Attitude     string `bson:"attitude"`
	Attachments  string `bson:"attachments"`
	Platform     string `bson:"platform"`
	MsgExtType   string `bson:"msg_ext_type"`
	AttacCount   int    `bson:"attac_count"`
	ContentLen   int    `bson:"content_len"`
	IsSaveToDb   int    `bson:"is_save_to_db"`
	MsgFlag      int    `bson:"msg_flag"`
	MsgStatus    int    `bson:"msg_status"`
	MsgType      int    `bson:"msg_type"`
	SendDate     int    `bson:"send_date"`
}
