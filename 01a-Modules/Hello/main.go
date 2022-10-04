package main

import (
	"fmt"

	"hakanaltindis/helper"
)

// go mod edit -replace example.com/greetings=../greetings

func main() {
	fmt.Println("Hakan!")

	helper.SayHi("Hakan", "Ahmet")
}
