package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Printf("hello, %d\n", i)
}

func main() {
	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go hello(i)
	}

	fmt.Println(0)
	wg.Wait()
	// main 函数之后，由 main 函数启动的 goroutine 也结束了
	time.Sleep(5 * time.Second)
}
