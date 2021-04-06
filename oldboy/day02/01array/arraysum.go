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
}
