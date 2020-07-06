//The break and continue keywords work just as they do in C and Java.
//A continue statement begins the next iteration of the innermost for loop at its post statement (i++).
package main

//keyword = for
import "fmt"

//Three-component loop

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}

	fmt.Println(sum)

}
