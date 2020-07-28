//if statement supports a composite syntax where the tested expression is preceded by an initialization statement

package main

import (
	"fmt"
)

func main() {
	if x := 100; x == 100 {
		fmt.Println("USA")
	}
}
