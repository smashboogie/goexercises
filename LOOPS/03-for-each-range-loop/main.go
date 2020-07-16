package main

//keyword = for
import "fmt"

//Looping over elements in slices, arrays, maps, channels or strings is often better done with a range loop.

func main() {
	strings := []string{"F", "Y"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
