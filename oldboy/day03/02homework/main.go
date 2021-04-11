package main

import "fmt"

/*
 * 有 50 枚金币，需要非配如下几人，分派规则：
 * a. 名字中每包含1个'e'或者'E'分 1 枚
 * b. 名字中每包含1个'i'或者'I'分 2 枚
 * c. 名字中每包含1个'o'或者'O'分 3 枚
 * d. 名字中每包含1个'u'或者'U'分 4 枚
 */

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emille", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth", "UU",
	}

	distribute = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Printf("left：%d\n", left)
}

func dispatchCoin() int {
	right_coin := 0
	for _, user := range users {
		for i := 0; i < len(user); i++ {
			switch user[i] {
			case 'e':
				distribute[user] += 1
			case 'E':
				distribute[user] += 1
			case 'i':
				distribute[user] += 2
			case 'I':
				distribute[user] += 2
			case 'o':
				distribute[user] += 3
			case 'O':
				distribute[user] += 3
			case 'u':
				distribute[user] += 4
			case 'U':
				distribute[user] += 4
			}
		}
		fmt.Printf("user: %s, coin: %d\n", user, distribute[user])
		right_coin += distribute[user]
	}

	left_coin := coins - right_coin
	if left_coin < 0 {
		fmt.Printf("there need %d coin, but only have %d\n", right_coin, coins)
		return 0
	}
	return left_coin
}
