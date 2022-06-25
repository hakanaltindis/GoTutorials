package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// See the below link, to more information about placeholders
// https://pkg.go.dev/fmt

func main() {

	// **** OUTPUT **** \\

	var (
		str1    string = "Ali"
		str2           = "topu"
		str3           = "tut."
		aNumber int    = 12
		isTrue  bool   = false
	)

	stringLength, _ := fmt.Println(str1, str2, str3)
	fmt.Println("String length:", stringLength)

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float32(aNumber))

	fmt.Printf("Data Types: %T, %T\n", aNumber, isTrue)

	// **** INPUT **** \\

	// Enter data to console

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n') // Attention, i did not use ":=" to assign

	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
