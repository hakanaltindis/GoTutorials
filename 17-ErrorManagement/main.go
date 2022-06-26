package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// 1. Handle error v1
	_, err := os.Open("dosyam.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 2. Handle error v2
	_, err = os.Open("dosyam.txt")
	checkError(err)

	// 3. Generate Error
	myError := errors.New("This is an error!")
	fmt.Println(myError.Error())

	// 4. return error with something else
	_, errResult := SendPositiveNumber(-5)
	fmt.Println(errResult)
}

// return error with something else
func SendPositiveNumber(number int) (int, error) {
	if number < 0 {
		return number, fmt.Errorf("The number has negative value!")
	}
	return number, nil
}

// Check error and if there is an error, exit the program
func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR :", err.Error())
		os.Exit(1)
	}
}
