package other

import "fmt"

// public function
func WritePublic(message string) {
	fmt.Println(message)
}

// private function
func writePrivate(message string) {
	fmt.Println(message)
}

// public const variable
const Pi float32 = 3.15

// private const variable
const pi float32 = 3.14
