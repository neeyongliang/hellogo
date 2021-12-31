package main

import (
	"fmt"
	"gpm/day02/error_handle"
	"gpm/day02/error_switch"
	"gpm/day02/fluent_interface"
	"gpm/day02/resource_clean"
)

func main() {
	fmt.Println("===>>> error switch example...")
	error_switch.Example()
	fmt.Println("===>>> resource clean example...")
	resource_clean.Example()
	fmt.Println("===>>> fluent interface example...")
	fluent_interface.Example()
	fmt.Println("===>>> error handle example...")
	error_handle.Example()
}
