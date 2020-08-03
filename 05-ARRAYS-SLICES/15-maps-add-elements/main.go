package main

import "fmt"

func main() {

	emps := make(map[string]int)

	emps["mark"] = 1
	emps["john"] = 2
	fmt.Println(emps)

	emps["sally"] = 3
	emps["sully"] = 4
	fmt.Println(emps)
}
