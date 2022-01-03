package main

import (
	"fmt"
	"gpm/day06/generate"
)

//go:generate ./gen.sh ./template/container.tmp.go generate uint32 container
func generateUint32Example() {
	var u uint32 = 42
	c := generate.NewUint32Container()
	c.Put(u)
	v := c.Get()
	fmt.Printf("generateExample: %d (%T)\n", v, v)
}

//go:generate ./gen.sh ./template/container.tmp.go generate string container
func generateStringExample() {
	var s string = "Hello"
	c := generate.NewStringContainer()
	c.Put(s)
	v := c.Get()
	fmt.Printf("generateExample: %s (%T)\n", v, v)
}

func main() {
	generateUint32Example()
	generateStringExample()
}
