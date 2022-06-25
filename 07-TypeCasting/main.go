package main

import (
	"fmt"
	"strconv"
)

func main() {

	// **** CONVERT **** \\
	var myString, floatString = "1", "2.145"
	var x = 10
	var pi float32 = 3.14

	fmt.Println(myString, x, pi)

	// convert to int from string
	number, _ := strconv.Atoi(myString)
	fmt.Println(number)

	// convert to string from int
	result := strconv.Itoa(x)
	fmt.Println("Result is " + result)

	// convert to float from string
	fl, _ := strconv.ParseFloat(floatString, 32)
	fmt.Println(fl)

	// convert to string from float
	s := fmt.Sprintf("%f", fl)
	fmt.Println(s)

	// **** CASTING **** \\

	var i int = 64
	var f = float32(i)
	var u = uint(f)

	// Do not permission implicit casting, you must explicit casting
	// var f2 float32 = i

	fmt.Println(i, f, u)
}
