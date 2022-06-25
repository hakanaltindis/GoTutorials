package main

import "fmt"

func main() {
	a := 10
	b := 10

	// class if clauses
	if b > a {
		fmt.Println("Büyüktür")
	} else if b == a {
		fmt.Println("Eşittir")
	} else {
		fmt.Println("Küçüktür")
	}

	// define variable in if scope
	if foo := 2; foo == 1 {
		fmt.Println("bar")
	} else {
		fmt.Println("buz")
	}
}
