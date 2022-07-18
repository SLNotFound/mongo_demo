package notice

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

var noticeDirArr = [2]string{"F:\\msg\\5A163DE9-731E-0072-8FCF-1F38E80D6B5F\\notice", "F:\\msg\\5A163DE9-731E-0072-8FCF-1F38E80D6B5F\\notice_mob"}

var f *os.File

func GetNtcFilePath(noticeDirArr [2]string) (ntcFilePathList []string) {
	for _, ntcPath := range noticeDirArr {

		ntcDirs, err := ioutil.ReadDir(ntcPath)
		if err != nil {
			fmt.Printf("notice file path read failed, err: %v\n", err)
			return nil
		}
		for _, ntcDir := range ntcDirs {
			if ntcDir.IsDir() {
				filepath := ntcPath + "\\" + ntcDir.Name()
				readDir, err := ioutil.ReadDir(filepath)
				if err != nil {
					log.Fatal(err)
				}
				for _, fileAllPath := range readDir {
					ntcFilePathList = append(ntcFilePathList, filepath+"\\"+fileAllPath.Name())
				}
			}
		}
	}
	return ntcFilePathList
}

func GetRecvId(ntcFilePath string) (recvId string) {
	var tempList []string
	tempList = strings.Split(ntcFilePath, "\\")
	return tempList[len(tempList)-2]

}

func SplitData(filePath string) (params []string, props map[string]string) {
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

	paramsData := strings.Split(dataList[0], " ")

	propsData := dataList[1:]

	var m = make(map[string]string)
	for _, v := range propsData {
		split := strings.Split(v, ":")
		if len(split) > 1 {
			fmt.Sprintf("%#v", split)
			m[split[0]] = strings.Trim(split[1], " ")
		}
	}

	return paramsData, m
}

func GetMsgId(msgFilePath string) string {
	split := strings.Split(msgFilePath, "\\")
	msgFile := split[len(split)-1]
	msgId := strings.Split(msgFile, ".")
	return msgId[0]
}

func ReadFromNte() (noticeList []*model.Notice) {

	filePathList := GetNtcFilePath(noticeDirArr)
	var notice *model.Notice
	for _, filePath := range filePathList {
		recvId := GetRecvId(filePath)
		msgId := GetMsgId(filePath)
		params, props := SplitData(filePath)
		method, _ := strconv.Atoi(params[0])
		createTime, _ := strconv.Atoi(params[len(params)-1])

		switch method {
		case 800:
			notice = &model.Notice{
				Method:     method,
				CreateTime: createTime,
				PcRead:     0,
				MobRead:    0,
				SendId:     params[2],
				RecvId:     recvId,
				MsgId:      msgId,
				Params:     params[1:],
				Props:      props,
			}
			noticeList = append(noticeList, notice)
		case 820:
			notice = &model.Notice{
				Method:     method,
				CreateTime: createTime,
				PcRead:     0,
				MobRead:    0,
				SendId:     params[2],
				RecvId:     recvId,
				MsgId:      msgId,
				Params:     params[1:],
				Props:      props,
			}
			noticeList = append(noticeList, notice)
		}
	}
	defer f.Close()
	return noticeList
}
