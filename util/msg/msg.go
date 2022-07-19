package msg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"mongo_demo/model"
	"os"
	"strconv"
	"strings"
)

var f *os.File

var noticeDir = "F:\\msg\\5A163DE9-731E-0072-8FCF-1F38E80D6B5F\\message"

func GetMsgFilePath(msgPath string) (msgFilePathList []string) {
	msgDirs, err := ioutil.ReadDir(msgPath)
	if err != nil {
		fmt.Printf("msg file path read failed, err: %v\n", err)
		return nil
	}

	for _, msgDir := range msgDirs {
		if msgDir.IsDir() {
			filepath := msgPath + "\\" + msgDir.Name()
			readDir, err := ioutil.ReadDir(filepath)
			if err != nil {
				log.Fatal(err)
			}
			for _, fileAllPath := range readDir {

				msgFilePathList = append(msgFilePathList, filepath+"\\"+fileAllPath.Name())
			}
		}
	}
	return msgFilePathList
}

func GetDate(msgFileDir string) string {
	var tempList []string
	tempList = strings.Split(msgFileDir, "\\")
	return tempList[len(tempList)-2]
}

func GetMsgId(msgFilePath string) string {
	split := strings.Split(msgFilePath, "\\")
	msgFile := split[len(split)-1]
	msgId := strings.Split(msgFile, ".")
	return msgId[0]
}

func SplitData(filePath string) (dataMap map[string]string) {
	var err error
	f, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var dataList []string
	for scanner.Scan() {
		dataList = append(dataList, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	InfoList := dataList[:9]
	content := dataList[len(dataList)-1]
	var m = make(map[string]string)
	for _, v := range InfoList {
		split := strings.Split(v, ":")
		if len(split) > 1 {
			fmt.Sprintf("%#v", split)
			m[split[0]] = strings.Trim(split[1], " ")
		}
	}
	m["content"] = content

	return m
}

func GetReceiver(dataMap map[string]string) (string, string) {
	receiver := dataMap["receive-users"]
	split := strings.Split(receiver, ";")
	return split[0], split[1]
}

func ReadDataFromMsg() (msgList []*model.Msg, date string) {
	msgFilePathList := GetMsgFilePath(noticeDir)
	var msg *model.Msg
	for _, msgFilePath := range msgFilePathList {
		date = GetDate(msgFilePath)
		dataMap := SplitData(msgFilePath)
		recvId, recvUser := GetReceiver(dataMap)
		msgId := GetMsgId(msgFilePath)
		sendDate, _ := strconv.Atoi(dataMap["senddate"])
		msg = &model.Msg{
			MsgId:        msgId,
			Subject:      dataMap["subject"],
			SendId:       dataMap["senderid"],
			SendName:     dataMap["sendername"],
			ReceiveId:    recvId,
			ReceiveName:  recvUser,
			ReceiveUsers: dataMap["receive-users"],
			SendIp:       "",
			SendMac:      "",
			DataPath:     date,
			ContentType:  dataMap["content"],
			ExtData:      "",
			SourceId:     "",
			Receiver:     recvUser,
			Attitude:     "",
			Attachments:  "",
			Platform:     "1",
			MsgExtType:   "",
			AttacCount:   0,
			ContentLen:   78,
			IsSaveToDb:   0,
			MsgFlag:      0,
			MsgStatus:    0,
			MsgType:      0,
			SendDate:     sendDate,
		}

		msgList = append(msgList, msg)
	}
	defer f.Close()
	return msgList, date
}
