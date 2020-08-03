package main

import "fmt"

func main() {

	emps := make(map[string]int)
	//adding elements
	emps["mark"] = 1
	emps["john"] = 2
	emps["sally"] = 3
	fmt.Println(emps)

	//updating an existing element
	emps["mark"] = 6
	fmt.Print(emps)

	//deleting an element
	delete(emps, "john")
	fmt.Println(emps)
}
