package main

import "fmt"

/*
 * A defer statement defers the execution of a function until the surrounding function returns.
 * The deferred call's arguments are evaluated immediately,
 * but the function call is not executed until the surrounding function returns.

 * Deferred function calls are pushed onto a stack. When a function returns,
 * its deferred calls are executed in last-in-first-out order LIFO
 */

var isConnected bool = false

func main() {

	// 1. Basic sample
	defer fmt.Println("world")
	fmt.Println("hello")

	// 2. Defer with parameters
	x, y := 3, 4
	defer add(x, y)
	x = 5
	add(x, y)

	// 3. Possible usage field
	fmt.Println()

}

func add(x, y int) {

	fmt.Println("Totatl :", x+y)
}

func databaseProcessing() {
	connect()
	defer disconnect()
	fmt.Println("Doing something")
}

func connect() {
	isConnected = true
	fmt.Println("Connected to database!")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected!")
}
