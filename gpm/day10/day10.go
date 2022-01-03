package main

import (
	"gpm/day10/simple_generics"
	"gpm/day10/type_generics"
)

// go run -gcflags=-G=3 ./main.go
func main() {
	simple_generics.PrintExample()
	simple_generics.FindExample()
	type_generics.StackExample()
	type_generics.ListExample()
}
