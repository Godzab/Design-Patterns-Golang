package main

// Dependency inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

type Relationship int

const(
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct{
	name string
}

type Info struct{
	from *Person
	relationship Relationship
	to *Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person){
	r.relations = append(r.relations, Info{parent, 1, child})
}
