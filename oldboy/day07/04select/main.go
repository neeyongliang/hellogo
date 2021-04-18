package main

import "fmt"

func main() {
	// 尝试修改通道的缓冲区大小，查看输出
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
