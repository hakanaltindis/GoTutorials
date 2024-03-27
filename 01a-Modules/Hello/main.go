package main

import (
	"fmt"

	"hakanaltindis/helper"
)

// go mod init example/hello
// go mod edit -replace example.com/greetings=../greetings
// go get hakan/helper
// go mod tidy

func main() {
	fmt.Println("Hakan!")

	helper.SayHi("Hakan", "Ahmet")
}
