package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	msg := "Hello"
	// use arguments
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
