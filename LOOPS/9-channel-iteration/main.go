//For channels, the iteration values are the
//successive values sent on the channel until closed.

package main

//keyword = for
import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
