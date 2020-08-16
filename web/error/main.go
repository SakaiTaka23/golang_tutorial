package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	f, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(f)
}

// Sqrt func
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// implementation
	f = math.Sqrt(f)
	return f, nil
}
