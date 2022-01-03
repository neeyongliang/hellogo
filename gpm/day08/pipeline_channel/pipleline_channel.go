package pipeline_channel

import "fmt"

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

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
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
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

type EchoFun func([]int) <-chan int
type PipeFun func(<-chan int) <-chan int

func pipeline(nums []int, echo EchoFun, pipeFns ...PipeFun) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func Example() {
	fmt.Println("===>> pipeline channel example...")
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println(n)
	}
}
