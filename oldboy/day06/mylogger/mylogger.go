package mylogger

import (
	"errors"
	"strings"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger logger struct
type Logger struct {
	level LogLevel
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("无效的日志级别错误")
	}

}

// NewLog new function
func NewLog(level string) Logger {
	l, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	return Logger{level: l}
}
