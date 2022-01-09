package exercises

import "fmt"

func gen(n int) []int {
	fmt.Println("feibonaqie example...")
	if n < 2 {
		return nil
	}
	// var fb = make([]int, 2)
	var fb = []int{}
	var fbMap = make(map[int]int)
	fbMap[0] = 1
	fbMap[1] = 1
	fb = append(fb, fbMap[0], fbMap[1])
	for i := 2; i < n; i++ {
		fbMap[i] = fbMap[i-1] + fbMap[i-2]
		fb = append(fb, fbMap[i])
	}
	return fb
}

func Feibonaqie(n int) {
	ret := gen(n)
	fmt.Println(ret)
}
