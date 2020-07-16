//The break and continue keywords work just as they do in C and Java.
//The break statement works the same with every loop we make.
//the break ends the loop when executed
package main

//keyword = for
import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i >= 7 {
			break
		}
		fmt.Print(i, "  ")
	}
}
