package main

import "fmt"

func binary_search(l []int, item int) (ret int) {
	var low = 0
	var height = len(l) - 1
	var mid = 0
	var guess = 0

	for low < height {
		fmt.Println("try binary search...")
		mid = (low + height) / 2
		guess = l[mid]
		if guess == item {
			return item
		}
		if guess > item {
			height = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	// var t := make([]int, 10)
	var t = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var ret = binary_search(t, 3)
	fmt.Println(ret)
}
