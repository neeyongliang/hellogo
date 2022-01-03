package main

import (
	"fmt"
	"gpm/day03/functional_builder"
	"gpm/day03/server_builder"
)

func main() {
	fmt.Printf("===>> server new, not recommend...\n")
	server_builder.TryNew("127.0.0.1", 9989, nil)
	fmt.Printf("===>> server builder, not recommend...\n")
	server_builder.TryBuilder()
	fmt.Printf("===>> server functional, recommend...\n")
	functional_builder.Example()
}
