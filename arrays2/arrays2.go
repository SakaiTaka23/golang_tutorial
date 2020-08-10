package main

import (
	"fmt"
)

func main() {
	fruitSlice := []string{"Apple","Orange","Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}