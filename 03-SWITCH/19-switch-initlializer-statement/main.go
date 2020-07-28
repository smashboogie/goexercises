//The switch keyword may be immediately followed by a simple initialization statement where variables,
//local to the switch code block, may be declared and initialized.

package main

import (
	"fmt"
	"time"
)

func main() {
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("today is less than 5th")
	case today.Day() == 10:
		fmt.Println("today is the 10th")
	case today.Day() > 15:
		fmt.Println("today is more than the 15th")
	default:
		fmt.Println("default case")
	}
}
