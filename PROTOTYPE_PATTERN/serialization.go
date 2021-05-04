package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//Knows how to unwrap a structure

type Address2 struct{
	StreetAddress, City, Country string
}

type Person2 struct{
	Name string
	Address *Address1
	Friends []string
}

func (p *Person2) DeepCopy()*Person2{
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d:= gob.NewDecoder(&b)
	res := Person2{}
	_ = d.Decode(&res)

	return &res
}

func main(){
 //Peform some copy functions
}
