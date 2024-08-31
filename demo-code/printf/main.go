package main

import "fmt"

func main() {
	print(Customer{
		ID:      1,
		Name:    "Foo",
		Country: "Bar",
	}.String())
}

// START // OMIT

type Customer struct {
	ID      int
	Name    string
	Country string
}

func (c Customer) String() string {
	return fmt.Sprintf("%s: %s (%s)", c.ID, c.Name, c.Country) // HL
}

// END // OMIT
