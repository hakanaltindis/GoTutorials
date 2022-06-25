package main

import "fmt"

func main() {

	// break
	i := 0

	for {
		if i >= 3 {
			break
		}

		fmt.Println("i'nin değeri :", i)
		i++
	}

	// continue
	for j := 0; j < 7; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println("Çıktı :", j)
	}
}
