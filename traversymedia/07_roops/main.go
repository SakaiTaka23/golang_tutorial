package main

import (
	"fmt"
)

func main() {
	// long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// fizz buzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d is fizz buzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d is fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d is buzz\n", i)
		} else {
			fmt.Printf("%d is no\n", i)
		}
	}
}
