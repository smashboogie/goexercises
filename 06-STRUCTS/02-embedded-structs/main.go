package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

type Manager struct {
	Employee
	isMgr bool
}

func main() {

	mgr1 := Manager{
		Employee: Employee{
			firstName: "Hank",
			lastName:  "Gathers",
			age:       30,
			salary:    200,
		},
		isMgr: true,
	}

	emp2 := Employee{"Sally", "Stothers", 40, 300}

	fmt.Println(mgr1)
	fmt.Println(emp2)
}
