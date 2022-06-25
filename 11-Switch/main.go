package main

import "fmt"

func main() {
	foo := 3

	// define classic switch
	switch foo {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("None")
	}

	// do not use condition in switch
	// condition in case statements
	switch {
	case foo == 3:
		fmt.Println("Three")
	case foo > 1:
		fmt.Println("One")
	default:
		fmt.Println("None")
	}

	// e.g Not required break
}
