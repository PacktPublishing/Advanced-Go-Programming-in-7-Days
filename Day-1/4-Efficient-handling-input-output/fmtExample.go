package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("%v (%d €)", p.Name, p.Price)
}

func main() {
	// Fmt relies on the Stringer interface to print custom types
	s := fmt.Sprint(Product{"Quirky Pants", 100})
	// build-in function
	println(s)
}
