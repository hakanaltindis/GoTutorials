package main

/*
 * First of all, run 'go mod init ...' command to track dependency for all folders.
 * Add sample code blocks to show up
 */

// go get other
import (
	"fmt"
	"other"
)

func main() {

	// can call only public function
	other.WritePublic("Public Message!")

	// can call only public const variable
	fmt.Println(other.Pi)
}
