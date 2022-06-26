package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// 1. Basic sample
	fmt.Println("Hello World")
	// To change value of random in Intn function
	rand.Seed(time.Now().UnixNano())

	fmt.Println("My number is", rand.Intn(10))

	fmt.Println(strings.Contains("Test", "st"))

	fmt.Println(strings.Count("test", "t"))

	fmt.Println(strings.HasPrefix("unit_test", "unit_"))

	fmt.Println(strings.HasSuffix("file.rar", ".rar"))

	fmt.Println(strings.Index("Hello World", "l"))

	// 2. Generate your package

	// Not completed
}
