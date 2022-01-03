package simple_generics

import "fmt"

// go run -gcflags=-G=3 xxx.go
// 函数名 约束条件 参数类型

func myprint[T any](arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func find[T comparable](arr []T, elem T) int {
	for i, v := range arr {
		if v == elem {
			return i
		}
	}
	return -1
}

func PrintExample() {
	fmt.Println("===>> simple generics print example...")
	strs := []string{"Hello", "World", "Generics"}
	decs := []float64{4.1, 2.9, 3.2, 2.1, 1.14}
	nums := []int{2, 4, 6, 8, 10}

	myprint(strs)
	myprint(decs)
	myprint(nums)
}

func FindExample() {
	fmt.Println("===>> simple generics find example...")
	strs := []string{"Hello", "World", "Generics"}
	decs := []float64{4.1, 2.9, 3.2, 2.1, 1.14}
	nums := []int{2, 4, 6, 8, 10}

	fmt.Printf("find World in strs: %d\n", find(strs, "World"))
	fmt.Printf("find 2.1 in decs: %d\n", find(decs, 2.9))
	fmt.Printf("find 4 in nums: %d\n", find(nums, 4))
}
