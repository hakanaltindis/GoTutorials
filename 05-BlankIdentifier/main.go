package main

import (
	"errors"
	"fmt"
)

func main() {
	// Blank identifier ( _ )
	// It is only used in Go Lang first time.

	// use when return multi values
	result, _ := checkBiggerThan10(11)
	fmt.Println(result)

	// other simple definitions
	var _, a = 10, "string"
	fmt.Println(a)

}

func checkBiggerThan10(number int) (bool, error) {

	if number > 10 {
		return true, nil
	}

	return false, errors.New("It is not bigger than 10")
}
