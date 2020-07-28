// Golang program to iterate over
// an Array using for loop
package main

import "fmt"

func main() {

	// taking an array
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("the array of the loop is: ")
	// using for loop
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
