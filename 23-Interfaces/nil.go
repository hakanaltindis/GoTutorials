package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	// t is typed nil
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	// runtime error
	var i2 I
	describe(i2)
	i2.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
