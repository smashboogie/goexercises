//For strings, the range loop iterates over Unicode code points.
//The rune type is an alias for int32, and is used to emphasize that an integer represents a code point.

package main

//keyword = for
import "fmt"

func main() {
	for i, l := range "DJKLDJDK$4784" {
		fmt.Printf("%#U is on the byte position: %d\n", l, i)
	}
}
