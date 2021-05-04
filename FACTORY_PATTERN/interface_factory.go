package main

import "fmt"

type person struct{
	name string
	age int
}

type oldPerson struct{
	name string
	age int
}


type Person1 interface {
	SayHello()
}

func (p *person) SayHello(){
	fmt.Printf("Hi, my name is %s. I am %d years old.\n", p.name, p.age)
}

func (p *oldPerson) SayHello(){
	fmt.Printf("Hi, my old ass name is %s. I am %d years old.\n", p.name, p.age)
}

func NewPersons(name string, age int)Person1{
	return &person{name, age}
}

func NewOldPersons(name string, age int)Person1{
	return &oldPerson{name, age}
}

func main(){
	p := NewPersons("James", 34)
	p.SayHello()

	op := NewOldPersons("Geeza", 74)
	op.SayHello()
}
