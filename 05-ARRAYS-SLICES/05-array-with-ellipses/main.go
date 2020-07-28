// use the ellipses to inver array length when using array literals
//we can use '...' to infer the length based on the elements inside {}

package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(len(a))
}
