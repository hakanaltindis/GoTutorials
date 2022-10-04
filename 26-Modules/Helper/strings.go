package helper

import "fmt"

func SayHi(strings ...string) {
	for i := 0; i < len(strings); i++ {
		sayHiPrivate(strings[i])
	}
}

func sayHiPrivate(name string) {
	fmt.Printf("Hi, %s\n", name)
}
