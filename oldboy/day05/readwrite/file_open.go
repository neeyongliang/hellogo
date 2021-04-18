package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const LOG_FILE string = "./main.go"

func readFromFile(name string) {
	fileObj, err := os.Open(name)
	if err != nil {
		fmt.Printf("Open file failed, %s", err)
		return
	}

	defer fileObj.Close()
	// var tmp = make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("Read complte！")
			return
		}

		if err != nil {
			fmt.Printf("Read file failed, err: %s", err)
			return
		}

		fmt.Printf("Get %d byte", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

func readFromIOBuffer(name string) {
	fileObj, err := os.Open(name)
	if err != nil {
		fmt.Printf("Open file failed, %s", err)
		return
	}

	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}

		fmt.Print(line)
	}
}

func readFromFileByUtil(name string) {
	ret, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	fmt.Print(string(ret))
}

func main() {
	readFromFile(LOG_FILE)
	fmt.Println("=========================")
	readFromIOBuffer(LOG_FILE)
	fmt.Println("=========================")
	readFromFileByUtil(LOG_FILE)
}
