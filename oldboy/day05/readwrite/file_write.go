package main

import "os"

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer fileObj.Close()
	fileObj.Write([]byte("this is string"))
	fileObj.WriteString("\nthat a int")

}

func writeDemo2() {
	// 类似的还有 bufio, ioutil.Writer
	// 针对 bufio 的操作是写到缓存里，需要 xx.Flush()
}

func main() {
	writeDemo1()
}
