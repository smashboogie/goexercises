//switch statement is used to select one of many blocks of code to be executed.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is the 5th")
	case 10:
		fmt.Println("Today is the 10th")
	default:
		fmt.Println("Default day")
	}
}
