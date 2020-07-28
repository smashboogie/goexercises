package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("the outer loop is %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tthe inner loop is %d\n", j)
		}
	}
}
