//The case statement can also used with conditional operators.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch {
	case today.Day() < 5:
		fmt.Println("today is less than 5th")
	case today.Day() == 10:
		fmt.Println("today is the 10th")
	case today.Day() > 15:
		fmt.Println("today is greater than 15th")
	default:
		fmt.Println("the default day")
	}
}
