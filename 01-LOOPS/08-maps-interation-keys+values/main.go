//iteration order over maps is not specified and
//is not guaranteed to be the same from one iteration to the next.

package main

import "fmt"

//keyword = for

func main() {
	m := map[string]int{
		"map1": 1,
		"map2": 2,
		"map3": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
