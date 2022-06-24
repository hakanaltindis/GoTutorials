package main

import "fmt"

// 9. define variable out of scope
var number = 5

/*
 * cannot define like this
 *  number := 2
 */

func main() {

	// 1. just define variable
	var message string
	message = "Hello, World!"
	fmt.Println(message)

	// 2. define variable with assinging first value
	var message2 string = "Hello, World 2!"
	fmt.Println(message2)

	// 3. define variable without specified type
	var message3 = "Hello, World 3!"
	fmt.Println(message3)

	// 4. define variable without type and var keyword
	message4 := "Hello, World 4!"
	fmt.Println(message4)

	// 5. define multiple variable
	var a, b, c int
	a = 3
	b = 4
	c = 5
	fmt.Println(a, b, c)

	// 6. define multiple variable with first value
	var a2, b2, c2 int = 3, 4, 5
	fmt.Println(a2, b2, c2)

	// 7. define multiple variable with mixed types
	var msg = "Hakan"
	var a3, b3, c3 = 2, false, 4.5
	fmt.Println(msg, a3, b3, c3)

	// 7. define multiple variable with the shortest way
	msg2, x, y := "Hakan", 5, true
	fmt.Println(msg2, x, y)

	// 8. define complex variable
	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)
	fmt.Println(myFloat32)
	fmt.Println(myComplex)

	// 9.
	fmt.Println(number)

	// 10. define const variable
	const constantNumber = 10

	// 11. define variable with scope
	var (
		num      = 1
		message5 = "Scope Message"
	)
	fmt.Println(num, message5)

}
