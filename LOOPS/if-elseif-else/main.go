package main //if...else if....else statement - executes different codes for more than two conditions

import (
	"fmt"
)

func main() {
	x := 100

	if x == 100 {
		fmt.Println("GER")
	} else if x == 50 {
		fmt.Println("JAPAN")
	} else {
		fmt.Println("USA")
	}
}
