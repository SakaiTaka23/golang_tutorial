package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	var students [3]string
	fmt.Printf("students: %v\n", students)
	students[0] = "Ahmed"
	students[1] = "Arnold"
	students[2] = "Lisa"
	fmt.Printf("students: %v\n", students)
	fmt.Printf("students0: %v\n", students[2])
	fmt.Printf("number of students: %v\n", len(students))

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2, 3}
	d := c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("length: %v\n", len(c))
	fmt.Printf("capacity: %v\n", cap(c))

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[:]
	g := e[3:]
	h := e[:6]
	i := e[3:6]
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

}
