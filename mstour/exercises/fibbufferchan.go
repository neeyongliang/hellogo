package exercises

import (
	"fmt"
	"time"
)

var quit = make(chan bool)

func fibnobuffer(c chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func FeibonaqieBufferChan() {
	start := time.Now()
	command := ""
	data := make(chan int)

	// 此处会等待输入
	go fibnobuffer(data)

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
