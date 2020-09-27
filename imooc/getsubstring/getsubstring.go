package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0
	fmt.Printf("origin string: %s\n", s)
	for i, ch := range []byte(s) {
		fmt.Printf("i: %d, ch: %c, start: %d\n", i, ch, start)
		if lastOccured[ch] >= start {
			start = lastOccured[ch] + 1
			fmt.Printf("bigger start, %c\n", lastOccured[ch])
		} else {
			fmt.Printf("less start, %c\n", lastOccured[ch])
		}

		if i-start+1 > maxLength {
			fmt.Printf("update max length: %d\n", maxLength)
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}

	return maxLength
}

func main() {
	l := lengthOfNonRepeatingSubStr("abcabcbb")
	fmt.Println(l)
}
