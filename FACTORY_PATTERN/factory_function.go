package main

type Person struct {
	Name string
	Age int
	EyeCount int
}

func NewPerson(name string, age int) *Person{
	if age < 16{
		// Do something here
	}
	 return &Person{name, age, 2}
}

func main(){
	p := NewPerson("Godfrey", 30)
	p.EyeCount = 3
}