package main

import "fmt"

// Color OCP
// Open for extension and closed for modification
// Specification (Enterprise pattern)
type Color int

const(
	red Color = iota
	green
	blue
)

type Size int

const(
	small Size = iota
	medium
	large
)

type Product struct{
	name string
	color Color
	size Size
}

type Filter struct{

}

// FilterByColor BAD way of filtering in an object
func (f Filter) FilterByColor(products []Product, color Color)[]*Product{
	result := make([]*Product, 0)
	for i, v := range products{
		if v.color == color{
			result = append(result, &products[i])
		}
	}
	return result
}

// Specification Pattern for better way of abstracting filters
type Specification interface {
	isSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) isSatisfied(p *Product) bool {
	return p.size == s.size
}

func (c ColorSpecification) isSatisfied(p *Product) bool {
	return p.color == c.color
}

type AndSpecification struct{
	first, second Specification
}

func (a AndSpecification) isSatisfied(p *Product) bool{
	return a.first.isSatisfied(p) && a.second.isSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification)[]*Product{
	result := make([]*Product, 0)
	for i, v := range products{
		if spec.isSatisfied(&v){
			result = append(result, &products[i])
		}
	}
	return result
}

func main(){
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, medium}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green){
		fmt.Printf("- %s is green\n", v.name)
	}

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	ft := BetterFilter{}
	for _, v := range ft.Filter(products, greenSpec){
		fmt.Printf("- %s is green\n", v.name)
	}
}
