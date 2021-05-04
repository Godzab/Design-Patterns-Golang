package main

import "fmt"

type Employee struct{
	Name, Position string
	AnnualIncome int
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee{
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct{
	Name, Position string
	AnnualIncome int
}

func main(){
	//Created a predefined grouping here which is nice
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 40000)

	developer:= developerFactory("Adam")
	manager:= managerFactory("Maggy")

	fmt.Println(developer, manager)
}