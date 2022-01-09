package exercises

import "fmt"

func FindPrimes() {
	fmt.Println("example find primes...")
	for i := 1; i <= 20; i++ {
		if func(number int) bool {
			if number <= 1 {
				return false
			}
			for j := 2; j < i; j++ {
				if number%j == 0 {
					return false
				}
			}
			return true
		}(i) {
			fmt.Println(i)
		}
	}
}
