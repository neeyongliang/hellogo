package main

import "fmt"

func main() {
	// 没有分配内存
	var s1 []int
	s1 = []int{1, 2, 3}
	fmt.Println(s1)

	// make初始化，分配内存
	s2 := make([]int, 3)
	fmt.Println(s2)

}
