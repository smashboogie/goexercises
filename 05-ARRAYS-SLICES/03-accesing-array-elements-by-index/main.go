//access or assign the array elements by referring to the index number
//The index is specified in square brackets.

package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "A"
	a[1] = "B"
	a[2] = "C"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a)
}
