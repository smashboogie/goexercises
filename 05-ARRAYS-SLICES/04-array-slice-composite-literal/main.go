// initialize an array with pre-defined values using an array literal
package main

import "fmt"

func main() {
	// Intialized with values
	x := [5]int{1, 2, 3, 4, 5}
	// Partial assignment
	var y [5]int = [5]int{6, 7, 8}
	fmt.Println(x)
	fmt.Println(y)
}
