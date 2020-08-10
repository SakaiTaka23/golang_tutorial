package main

import (
	"fmt"
)

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// loop through ids
	// range returns index,id
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

	for _, v := range emails {
		fmt.Println("Values: " + v)
	}
}
