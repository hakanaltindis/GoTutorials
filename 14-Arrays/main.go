package main

import "fmt"

/*
 * An array's length is part of its type, so arrays cannot be resized.
 * This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
 */

func main() {

	// 1. Basic array sample
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 2. Assign value on initialization time
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 3. Manipulate items of array
	var colors = [...]string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	fmt.Println(colors[1])

	colors[1] = "Yellow"
	fmt.Println(colors[1])

	// 4. Length of Array
	length := len(primes)
	fmt.Println("Number of Primes :", length)
	fmt.Println("Number of Colors :", len(colors))

	// 5. Capacity of Array
	fmt.Println("Capacity of Primes :", cap(primes))

	// 6. iterate item with for
	i := 0
	for i < len(colors) {
		fmt.Println(colors[i])
		i++
	}
}
