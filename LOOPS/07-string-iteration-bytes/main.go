//For strings, the range loop iterates over Unicode code points.
//To loop over individual bytes, simply use a normal for loop and string indexing

package main

//keyword = for
import "fmt"

func main() {
	const s = "fjkfjlkj"
	for i := 1; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

}
