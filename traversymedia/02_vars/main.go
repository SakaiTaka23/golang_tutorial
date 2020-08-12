package main

import (
	"fmt"
)

func main() {
	// var name string = "Brad"
	//var name = "Brad"

	// shorthand
	// name := "Brad"
	// email := "brad@gmail.com"

	name, email := "Brad", "brad@gmail.com"
	name, email = "rad","rad@gmail.com"

	fmt.Println(name, email)
	fmt.Printf("%T\n", name)
}
