package main

import (
	"fmt"
	"io"
	"os"
)

/*
 * A defer statement defers the execution of a function until the surrounding function returns.
 * The deferred call's arguments are evaluated immediately,
 * but the function call is not executed until the surrounding function returns.

 * Deferred function calls are pushed onto a stack. When a function returns,
 * its deferred calls are executed in last-in-first-out order LIFO
 */

/*
  * ***** RULES OF Defer *****
	* 1. A deferred function’s arguments are evaluated when the defer statement is evaluated.
	* 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	* 3. Deferred functions may read and assign to the returning function’s named return values.
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

/*
 *   || Excellent Sample for defer usage	||
 *   ||																		||
 *   \/																		\/
 */
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

/*
 * This works, but there is a bug. If the call to os.Create fails, the function will return without closing the source file.
 * This can be easily remedied by putting a call to src.Close before the second return statement,
 * but if the function were more complex the problem might not be so easily noticed and resolved.
 * By introducing defer statements we can ensure that the files are always closed:
 */

func CopyFile2(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// Other good example
func c() (i int) {
	defer func() { i++ }()
	return 1
}
