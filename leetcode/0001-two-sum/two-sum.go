package two_sum

import "fmt"

func twoSum1(nums []int, target int) []int {
	rawLen := len(nums)
	if rawLen < 2 {
		panic("less number!")
	}
	var newArray = []int{0, 0}
	for i := range nums {
		for j := i + 1; j <= rawLen-1; j++ {
			if nums[i]+nums[j] == target {
				newArray[0] = i
				newArray[1] = j
				break
			}
		}
	}
	return newArray
}

func twoSum2(nums []int, target int) []int {
	rawLen := len(nums)
	if rawLen < 2 {
		panic("nums too few...")
	}
	// map[target - value] = index
	newMap := make(map[int]int)
	// i 是当前位置，wanted 是回头看的配对位置
	for i, v := range nums {
		if wanted, ok := newMap[v]; ok {
			return []int{wanted, i}
		}
		newMap[target-v] = i
		fmt.Println(newMap)
	}
	return nil
}

// func main() {
// 	nnn := []int{2, 5, 7, 9, 11}
// 	s := twoSum1(nnn, 9)
// 	fmt.Printf("===>> twoSum1: %v\n", s)
// 	s = twoSum2(nnn, 9)
// 	fmt.Printf("===>> twoSum2: %v", s)
// }
