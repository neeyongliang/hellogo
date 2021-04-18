package main

import "fmt"

// 有 n 级台阶，一次可以走一步，一次也可以走两步，共有多少次走法
func upFloor(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 当最后只剩两层台阶时候，有以下走法：
	// 一次性迈过
	// 分两次迈过
	return upFloor(n-1) + upFloor(n-2)
}

func jieCheng(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * jieCheng(n-1)
}

func main() {
	res := jieCheng(5)
	fmt.Println("jie cheng...", res)

	ret := upFloor(4)
	fmt.Println("shang tai jie...", ret)

}
