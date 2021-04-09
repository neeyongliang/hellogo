package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1. 查找字符串中，汉字的数量
	var s1 = "hello上海"
	var count = 0
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	// 2. 计算出 how do you do 每个单词的数量
	var s2 = "how do you do"
	var s3 = strings.Split(s2, " ")
	word_counter := make(map[string]int, 10)
	for _, item := range s3 {
		if _, ok := word_counter[item]; !ok {
			word_counter[item] = 1
		} else {
			word_counter[item]++
		}
	}

	for key, value := range word_counter {
		fmt.Println(key, value)
	}

	// 3. 回文判断
	ss := "上海自来水来自海上"
	r := make([]rune, 0, len(ss))
	for i := 0; i < len(r)/2; i++ {
		// s[i] => s[len(s) - 1 - i]
		if r[i] != r[len(r)-1-i] {
			return
		}
	}
	fmt.Printf("ss: %s 是回文", ss)
}
