package main

import "fmt"

type stringfy interface {
	String() string
}

func Concat[T stringfy](vals []T) string {
	result := ""
	for _, v := range vals {
		result += v.String() + "\n"
	}
	return result
}

type Human struct {
	FirstName string
	LastName  string
}

func (h *Human) String() string {
	return fmt.Sprintf("First Name: %v, Last Name: %v", h.FirstName, h.LastName)
}

type Car struct {
	Brand string
	Model string
}

func (c *Car) String() string {
	return fmt.Sprintf("Brand: %v, Model: %v", c.Brand, c.Model)
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	h1 := new(Human)
	h1.FirstName = "Hakan"
	h1.LastName = "Altındiş"

	c1 := new(Car)
	c1.Brand = "Ford"
	c1.Model = "Mustang"

	var arr = []stringfy{h1, c1}
	fmt.Println(Concat(arr))
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func Contains[T comparable](s []T, x T) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}
