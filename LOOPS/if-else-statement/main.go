package main //if...else statement - executes some code if a condition is true and another code if that condition is false

import (
	"fmt"
)

func main() {
	x := 10

	if x == 100 {
		fmt.Println("GREAT")
	} else {
		fmt.Println("NOT GREAT")
	}
}
