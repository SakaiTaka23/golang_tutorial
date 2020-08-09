package main

import(
	"fmt"
)

const (
	_ = iota +5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main(){
	fmt.Printf("%v\n",catSpecialist)
}