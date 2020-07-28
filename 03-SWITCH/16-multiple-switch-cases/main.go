//he switch with multiple case line statement is used to select common block of code for many similar cases.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	var t int = today.Day()

	switch t {
	case 5, 10:
		fmt.Println("today is 5th or 10th")
	case 11, 12:
		fmt.Println("today is 11th or 12th")
	default:
		fmt.Println("today is default")
	}
}
