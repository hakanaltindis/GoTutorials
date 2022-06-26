package main

import "fmt"

func main() {

	// basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value :", i)
	}

	// There is no while keyword in GoLang
	// while loop with for keyword
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println("Sum :", sum)
	}

	// The range form of the for loop iterates over a slice or map
	// for-range keyword
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	a := [...]string{"a", "b", "c", "d", "e"}
	for i := range a {
		fmt.Println("Array item ", i, "is", a[i])
	}

	// for-range over map
	capitals := map[string]string{"Turkey": "Ankara", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}
