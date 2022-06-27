package main

/*
 * Go does not have classes. However, you can define methods on types.
 * A method is a function with a special receiver argument.
 * The receiver appears in its own argument list between the func keyword and the method name.
 * In this example, the Abs method has a receiver of type Vertex named v.

 * You can just define method for a type in same package!!!!!!
 */

import (
	"fmt"
	"math"
)

// 1. Define method for a struct
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 2. Define method for a custom type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// * You cannot define method for a type in another package!
// func (i int) Abs() int {
// 	if i < 0 {
// 		return int(-i)
// 	}
// 	return int(i)
// }

// 3. Pointer receiver to manipulate fields in strcut
// * For the statement v.Scale(5), even though v is a value and not a pointer,
// * the method with the pointer receiver is called automatically.
// * That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since
// * the Scale method has a pointer receiver.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// ***** Choosing a value or pointer receiver ***** \\
/*
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
*/
