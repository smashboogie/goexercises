package main

import "fmt"

func main() {
	anon := struct {
		firstName, lastName string
		age                 int
		salary              int
	}{
		firstName: "Sally",
		lastName:  "Jones",
		age:       30,
		salary:    120,
	}
	fmt.Println(anon)
}
