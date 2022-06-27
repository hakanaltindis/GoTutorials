package main

import (
	"fmt"
)

func main() {
	Alphabet()

	// // goroutine
	// go Alphabet()
	// time.Sleep(100 * time.Millisecond)

	// // goroutine
	// runtime.GOMAXPROCS(1)
	// go Alphabet2()
	// time.Sleep(100 * time.Millisecond)
}

func Alphabet() {
	for l := byte('a'); l <= byte('z'); l++ {
		fmt.Println(string(l))
	}
}

func Alphabet2() {
	for l := byte('a'); l <= byte('z'); l++ {
		go fmt.Println(string(l))
	}
}
