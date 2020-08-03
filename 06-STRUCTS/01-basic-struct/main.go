package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp1 := Employee{
		firstName: "mark",
		lastName:  "jones",
		age:       30,
		salary:    120,
	}
	emp2 := Employee{"Sally", "Stothers", 40, 300}

	fmt.Println(emp1)
	fmt.Println(emp2)
}
