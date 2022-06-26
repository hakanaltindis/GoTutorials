package main

import (
	"fmt"
	"strings"
)

/*
 * An array has a fixed size. A slice, on the other hand,
 * is a dynamically-sized, flexible view into the elements of an array.
 * In practice, slices are much more common than arrays.

 * Slices are like references to arrays
 * A slice does not store any data, it just describes a section of an underlying array.
 * Changing the elements of a slice modifies the corresponding elements of its underlying array.
 * Other slices that share the same underlying array will see those changes.
 */

func main() {

	// 1. Basic sample
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // indice of 1 is included, indice of 4 is not included
	fmt.Println(s)

	// 2. Slice defaults
	// * When slicing, you may omit the high or low bounds to use their defaults instead.
	// * The default is zero for the low bound and the length of the slice for the high bound.
	fmt.Println(primes[0:6])
	fmt.Println(primes[:6])
	fmt.Println(primes[0:])
	fmt.Println(primes[:])

	// 3. append
	fmt.Println("Before append, the slice is", s)
	s = append(s, 17)
	fmt.Printf("The slice is %v. Len of slice is %d. Cap of slice is %d.\n", s, len(s), cap(s))
	fmt.Println(primes)

	// 4. make and check the capacity
	numbers := make([]int, 5, 6)
	fmt.Println("The Numbers are", numbers, "The Length of nubmers is", len(numbers), "The capacity of numbers is", cap(numbers))
	numbers = append(numbers, 1)
	fmt.Println("The Numbers are", numbers, "The Length of nubmers is", len(numbers), "The capacity of numbers is", cap(numbers))
	numbers = append(numbers, 2)
	fmt.Println("The Numbers are", numbers, "The Length of nubmers is", len(numbers), "The capacity of numbers is", cap(numbers))

	// 5. Slice of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
