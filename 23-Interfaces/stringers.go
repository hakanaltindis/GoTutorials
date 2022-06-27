package main

/*
	One of the most ubiquitous interfaces is Stringer defined by the fmt package.

	type Stringer interface {
			String() string
	}
	A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/

import "fmt"

type Person struct {
	Name string
	Age  int
}

// like override ToString() method in C#
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Hakan Altındiş", 42}
	z := Person{"Mehmet Delibudak", 9001}
	fmt.Println(a, z)
}
