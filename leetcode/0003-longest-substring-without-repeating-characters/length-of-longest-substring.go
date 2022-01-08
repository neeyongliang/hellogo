package main

import "fmt"

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := map[byte]int{}

	end, maxLen := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for end+1 < n && m[s[end+1]] == 0 {
			// 没有出现在 map 集合中的数字是 0， 否则置为 1
			m[s[end+1]] = m[s[end+1]] + 1
			end++
			fmt.Println(m)
		}
		fmt.Println(end, i)
		fmt.Println(m)
		maxLen = max(maxLen, end-i+1)
	}
	return maxLen
}

func main() {
	var s string = "abcdsfgsa"
	fmt.Println(lengthOfLongestSubstring(s))
}
