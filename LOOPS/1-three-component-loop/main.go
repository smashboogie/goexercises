//https://yourbasic.org/golang/for-loop/

package main

//keyword = for
import "fmt"

//Three-component loop

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}
