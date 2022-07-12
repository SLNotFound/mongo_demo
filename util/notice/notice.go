package notice

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetNtcFilePath(ntcPath string) (ntcFilePathList []string) {
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
	return ntcFilePathList
}

func GetRecvId(ntcFilePath string) (recvId string) {
	var tempList []string
	tempList = strings.Split(ntcFilePath, "\\")
	return tempList[len(tempList)-2]

}

//func GetRecvId(ntcFilePathList []string) (recvIdList []string) {
//	var tempList []string
//	for _, ntcFilePath := range ntcFilePathList {
//		tempList = strings.Split(ntcFilePath, "\\")
//		recvId := tempList[len(tempList)-1]
//		recvIdList = append(recvIdList, recvId)
//	}
//
//	return recvIdList
//}
