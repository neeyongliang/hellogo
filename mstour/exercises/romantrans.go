package exercises

import (
	"fmt"
)

func RomanTrans(romans string) {
	transTable := map[rune]int{
		'M': 1000,
		'D': 5000,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	valcabular := make([]int, len(romans)+1)
	for index, digital := range romans {
		if val, ok := transTable[digital]; ok {
			valcabular[index] = val
		} else {
			fmt.Println("error, find illeagal character")
			return
		}
	}

	total := 0
	fmt.Println(valcabular)
	for i := 0; i < len(romans); i++ {
		if valcabular[i] < valcabular[i+1] {
			valcabular[i] = -valcabular[i]
		}
		total += valcabular[i]
	}

	fmt.Println(total)
}
