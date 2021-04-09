package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9}
	sum := 0

	for _, v := range a1 {
		sum += v
	}

	fmt.Printf("a1 sum: %d\n", sum)
	fmt.Println("=============================")

	var s1 map[string]int
	s1 = make(map[string]int, 2)
	s1["A"] = 19
	s1["B"] = 20
	v, ok := s1["C"]
	if !ok {
		fmt.Println("cannot find value")
	} else {
		fmt.Println(v)
	}
	fmt.Println("=============================")

	var ages [30]int
	ages = [30]int{1, 2, 3, 4, 5, 6}
	fmt.Println(ages)
	var age2 = [...]int{1, 2, 3, 4}
	fmt.Println(age2)
	var age3 = [...]int{1: 80, 90: 2}
	fmt.Println(age3)

	var a2 = [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a2)
}
