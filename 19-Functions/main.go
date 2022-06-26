package main

import "fmt"

func main() {

	// 1. just call function
	sayHello()

	// 2. pass parameter
	sayHelloToName("Hakan")

	// 3. return value
	var total = add(4, 5)
	fmt.Println(total)

	// 4. pass pointer as parameter
	sum := 0
	addWithSum(3, 4, &sum)
	fmt.Println(sum)

	// 5. multi return
	x, y := "Hello", "World"
	fmt.Println(x, y)
	x, y = swap(x, y)
	fmt.Println(x, y)

	// 6. varity parameter functions
	fmt.Println("Total of Terms :", addTerms(5, 10, 2, 3))

	// 7. Named return values
	fmt.Println(split(17))

	// 8. Anonymous functions
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, v := range terms {
			sum += v
		}
		numTerms = len(terms)
		return
	}

	fmt.Println(addFunc(2, 5, 7, 9))

}

func sayHello() {
	fmt.Println("Hello, guys")
}

func sayHelloToName(n string) {
	fmt.Println("Hello,", n)
}

func add(x, y int) int {
	return x + y
}

func addWithSum(x, y int, sum *int) {
	*sum = x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func addTerms(terms ...int) int {
	sum := 0
	for _, val := range terms {
		sum += val
	}
	return sum
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
