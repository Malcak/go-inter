package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "name"
	ftEmployee.age = 25
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	// GetMessage(ftEmployee) no se puede hacer de forma tan directa
}
