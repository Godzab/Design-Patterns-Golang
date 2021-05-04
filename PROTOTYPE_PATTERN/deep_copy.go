package main

import "fmt"

type Address struct{
	StreetAddress, City, Country string
}

type Person struct{
	Name string
	Address *Address
}

func main(){
	john := Person{"John",
		&Address{"D2 Avonsands Cresent", "Capetown", "South Africa"}}

	//This is an error prone statement
	Jane := john
	Jane.Name = "Jane" // This is okay
	Jane.Address.StreetAddress = "341 Baker street" // this is not okay because this is a pointer

	//This is the correct approach
	Jill := Jane
	Jill.Address = &Address{Jane.Address.StreetAddress, Jane.Address.City, Jane.Address.Country}
	Jill.Address.StreetAddress = "Sothwark Mews"
	fmt.Println(john, john.Address)
	fmt.Println(Jane, Jane.Address)
	fmt.Println(Jill, Jill.Address)
}