package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini: "port"`
	Username string `ini: "username"`
	Password string `ini: "password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Password string `ini:"password"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data parameter should be pointer")
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data parameter is not struct")
		return
	}

	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		err = errors.New("read config file error")
		return
	}

	lineSlice := strings.Split(string(b), "\n")
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		// comment
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// section
		if strings.HasPrefix(line, "[") {
			if line[len(line)-1] != ']' {
				err = fmt.Errorf("find section error in %d", idx)
				return
			}
			if len(strings.TrimSpace(line[1:len(line)-1])) == 0 {
				err = fmt.Errorf("section syntax error in %d", idx)
				return
			}
			fmt.Println(line)
		}
	}

	return nil
}

func main() {
	var mc MysqlConfig
	loadIni("./conf.ini", &mc)
	fmt.Println(mc.Address)
}
