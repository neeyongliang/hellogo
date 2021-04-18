package mylogger

// write log to console

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

func (l Logger) enable(logLevel LogLevel) bool {
	return l.level <= logLevel
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, a...)
	}
}

func (l Logger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a...)
	}
}

func (l Logger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a...)
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}

	funcName = runtime.FuncForPC(pc).Name() // xx包xx函数
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}
