package exercises

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, ch chan string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func FeibonaqieChan() {
	start := time.Now()

	size := 15

	ch := make(chan string, 15)
	for i := 0; i < size; i++ {
		go fib(float64(i), ch)
	}

	for i := 0; i < size; i++ {
		fmt.Printf(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
