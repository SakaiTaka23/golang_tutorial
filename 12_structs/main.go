package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//put same type in one row
	firstName, lastName, city, gender string
	age                               int
}

// greeting method value reciever(for just output)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method pointer reciever(for changeing values)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried pointer reciever
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative
	person2 := Person{"Bob", "Johnson", "NewYork", "m", 30}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
