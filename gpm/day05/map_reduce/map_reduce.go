package map_reduce

import (
	"fmt"
	"strings"
)

// map 是对批量数据进行同一种处理，所以输入与输出个数相同
func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapExample() {
	fmt.Println("===>> map example...")
	var list = []string{"Hao", "Chen", "MegaEase"}
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})

	fmt.Printf("%v\n", y)
}

// Reduce 是对一组数据进行连续处理
func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func ReduceExample() {
	fmt.Println("===>> reduce example...")
	var list = []string{"Hao", "Chen", "MegaEase"}
	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x)
}

func Filter(arr []int, fn func(n int) bool) []int {
	var newArray = []int{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func FilterExample() {
	fmt.Println("===>> filter example...")
	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})

	fmt.Printf("%v\n", out)
}
