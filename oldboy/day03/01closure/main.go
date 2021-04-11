package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// 闭包就是函数引用了作用域之外的变量，此处 suffix 就是外部变量
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	pngFunc := makeSuffixFunc(".png")
	fmt.Println(jpgFunc("test"))
	fmt.Println(pngFunc("test"))
	fmt.Println(pngFunc("aaa.png"))
}
