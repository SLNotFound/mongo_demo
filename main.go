package main

import (
	"fmt"
	"mongo_demo/util/notice"
	"os"
	"os/signal"
	"time"
)

func main() {
	start := time.Now()
	notice.ReadFromNte()
	end := time.Since(start)
	fmt.Println(end)
	var flag = make(chan os.Signal)
	signal.Notify(flag, os.Kill, os.Interrupt)
	<-flag
}
