package model

type Msg struct {
	MsgId        string
	Subject      string
	SendId       string
	SendName     string
	ReceiveId    string
	ReceiveName  string
	ReceiveUsers string
	SendIp       string
	SendMac      string
	DataPath     string
	ContentType  string
	ExtData      string
	SourceId     string
	Receiver     string
	Attitude     string
	Attachments  string
	Platform     string
	MsgExtType   string
	AttacCount   int
	ContentLen   int
	IsSaveToDb   int
	MsgFlag      int
	MsgStatus    int
	MsgType      int
	SendDate     int64
}
