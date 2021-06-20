package main

import "fmt"

func find_samllest(l []int) int {
	samllest := l[0]
	samllest_index := 0

	for i, _ := range l {
		if l[i] < samllest {
			samllest = l[i]
			samllest_index = i
		}
	}

	return samllest_index
}

func main() {
	l := []int{3, 5, 9, 1, 4, 2}
	s := []int{}
	samllest_index := -1
	// ret := find_samllest(l)
	// fmt.Println(ret)
	// new_array := make([]int, 10)
	var new_array []int
	for i, _ := range l {
		samllest_index = find_samllest(l)
		fmt.Println(l[samllest_index])
		new_array = append(new_array, l[samllest_index])
		s = append(l[:i], l[i+1:])
		l = s
	}

}
