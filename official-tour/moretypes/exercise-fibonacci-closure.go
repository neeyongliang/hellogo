package moretypes

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	pre := 0
	now := 1
	return func() int {
		now = pre + now
		pre = now - pre
		return pre
	}
}

func ExerciseFabonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

