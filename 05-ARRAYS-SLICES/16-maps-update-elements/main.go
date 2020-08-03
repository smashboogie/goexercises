package main

import "fmt"

func main() {

	emps := make(map[string]int)
	//adding elements
	emps["mark"] = 1
	emps["john"] = 2
	fmt.Println(emps)

	//updating an existing element
	emps["mark"] = 0
	fmt.Print(emps)
}
