package faninout

import (
	"fmt"
	"math"
	"sync"
)

// 使用 goroutine 和 channel，扇形输入输出

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sums = 0
		for i := range in {
			sums += i
		}
		out <- sums
		close(out)
	}()
	return out
}

// 传入 <-chan int 数组，合并为一个 <-chan int
func merge(cs []<-chan int) <-chan int {
	// 这里用到了 WaitGroup
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		// 取出所有的 channel
		go func(c <-chan int) {
			for n := range c {
				// 取出所有 channel 中的元素
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Example() {
	fmt.Println("===>> fan in out example...")
	nums := makeRange(1, 10000)
	in := echo(nums)

	// 使用 5 个 channel 分别计算质数之和
	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum(prime(in))
	}

	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
}
