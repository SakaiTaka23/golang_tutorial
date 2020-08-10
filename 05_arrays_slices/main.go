package main

import (
	"fmt"
)

func main() {
	// // Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))

	// starts from 1 stops at 2
	fmt.Println(fruitSlice[0:1])
	fmt.Println(fruitArr[1:2])
}
