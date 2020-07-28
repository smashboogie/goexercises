package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 7, 8)
	fmt.Println(x)
	y := []int{10, 15}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
