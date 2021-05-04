package main

import "fmt"

type Emp struct{
	Name, Position string
	AnnualIncome int
}

const(
	Developer = iota
	Manager
)

func NewEmployee(role int) *Emp{
	switch role {
	case Developer:
		return &Emp{"", "Developer", 60000}
	case Manager:
		return &Emp{"", "Manager", 40000}
	default:
		panic("Unsupported role")
	}
}

func main(){
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
