package main

import "fmt"

/*
 * A struct is a collection of fields.
 */

// Define struct
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Constructor is not support in GoLang
// Default function as constructor
func NewHuman() *Human {
	h := new(Human)
	return h
}

// No overloading in GoLang
func NewHumanWithFirstValues(f, l string, a int) *Human {
	h := new(Human)
	h.FirstName = f
	h.LastName = l
	h.Age = a
	return h
}

func main() {

	// 1. Generate struct
	h := Human{"Hakan", "Altındiş", 32}
	fmt.Println(h)

	// 2. Generate struct with new
	hNew := new(Human)
	hNew.FirstName = "Hakan"
	fmt.Println(hNew.FirstName)

	// 3. Generate struct
	h1 := Human{FirstName: "Hakan"} // assign default values to other fields
	fmt.Println(h1)

	// 4. Generate struct
	h2 := Human{} // assign default values to all fields
	fmt.Println(h2)

	// 4. Generate struct
	h3 := &Human{"Hakan", "Altındiş", 32} // assign default values to all fields
	fmt.Println(h3)

}
