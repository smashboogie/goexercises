//The fallthrough keyword used to force the execution flow to fall through the successive case block.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {

	case 5:
		fmt.Println("today 5th")
		fallthrough
	case 20:
		fmt.Println("today 20th")
		fallthrough
	case 15:
		fmt.Println("today 15th")
		fallthrough
	default:
		fmt.Println("today default")
	}
}
