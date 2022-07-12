package model

type Notice struct {
	Method  int
	PcRead  int
	MobRead int
	SendId  string
	RecvId  string
	MsgId   string
	Params  []string

	Props map[string]string
}