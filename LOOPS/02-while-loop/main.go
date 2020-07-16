package main

//keyword = for
import "fmt"

//while loop
//If you skip the init and post statements, you get a while loop.

func main() {
	n := 1
	for n < 5 {
		n *= 2
	}
	// 8 (1*2*2*2)
	fmt.Println(n)
}
