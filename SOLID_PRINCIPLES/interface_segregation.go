package main

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(document Document)
}

type multiFunctionalPrinter struct {}

func (m multiFunctionalPrinter) Print(d Document) {
	panic("implement me")
}

func (m multiFunctionalPrinter) Fax(d Document) {
	panic("implement me")
}

func (m multiFunctionalPrinter) Scan(document Document) {
	panic("implement me")
}

type oldFunctionPrinter struct{}

func (o oldFunctionPrinter) Print(d Document) {
	// WORKS
}

// Deprecated: Old fashioned printer does not need this
func (o oldFunctionPrinter) Fax(d Document) {
	panic("implement me")
}

// Deprecated: Old fashioned printer does not need this
func (o oldFunctionPrinter) Scan(document Document) {
	panic("implement me")
}
