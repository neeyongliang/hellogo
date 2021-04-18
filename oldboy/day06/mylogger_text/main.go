package main

import (
	"gogogo/oldboy/day06/mylogger"
	"time"
)

func main() {
	nnn := "2333333....."
	for {
		log := mylogger.NewLog("debug")
		log.Debug("这是一个 Debug 日志, %s", nnn)
		log.Warning("这是一个 Warning 日志")
		log.Info("这是一个 Info 日志")
		log.Error("这是一个 Error 日志")
		log.Fatal("这是一个 Fatal 日志")
		time.Sleep(1 * time.Second)
	}
}
