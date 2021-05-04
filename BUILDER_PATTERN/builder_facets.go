package main

import "fmt"

type Person struct{
	// Address
	StreetAddress, Postcode, City string

	//Job
	CompanyName, Position string
	AnnualIncome int
}

type PersonBuilder struct{
	person *Person
}

func(b *PersonBuilder) Lives()*PersonAddressBuilder{
	return &PersonAddressBuilder{*b}
}

func(b *PersonBuilder) Works()*PersonJobBuilder{
	return &PersonJobBuilder{*b}
}

func(b PersonBuilder) Build()*Person{
	return b.person
}

func NewPersonBuilder()*PersonBuilder{
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct{
	PersonBuilder
}

type PersonJobBuilder struct{
	PersonBuilder
}

func(it *PersonJobBuilder) AsA(position string)*PersonJobBuilder{
	it.person.Position = position
	return it
}

func(it *PersonJobBuilder) Earning(en int)*PersonJobBuilder{
	it.person.AnnualIncome = en
	return it
}

func(it *PersonJobBuilder) At(at string)*PersonJobBuilder{
	it.person.CompanyName = at
	return it
}

func(it *PersonAddressBuilder) In(city string)*PersonAddressBuilder{
	it.person.City = city
	return it
}

func(it *PersonAddressBuilder) WithPostcode(pc string)*PersonAddressBuilder{
	it.person.Postcode = pc
	return it
}

func(it *PersonAddressBuilder) At(at string)*PersonAddressBuilder{
	it.person.StreetAddress = at
	return it
}

func main(){
	pb := NewPersonBuilder()
	pb.
		Lives().
			At("123 London Road").
			In("London").
			WithPostcode("10234").
		Works().
			At("Double Eye").
			AsA("Developer").
			Earning(120000)
	person := pb.Build()
	fmt.Println(person)
}
