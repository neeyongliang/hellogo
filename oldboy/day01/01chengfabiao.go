package main

import "fmt"

func main() {
	var i int
	var j int
	for i = 1; i < 10; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%2d x%2d = %2d ", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println("=============================")

	for i = 9; i >= 1; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf("%2d x%2d = %2d ", i, j, i*j)
		}
		fmt.Println()
	}
}
