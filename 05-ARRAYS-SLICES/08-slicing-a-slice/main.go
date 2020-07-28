package main

import (
	"fmt"
)

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[4:])
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
