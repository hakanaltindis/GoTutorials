package main

import "fmt"

/*
 * A switch statement is a shorter way to write a sequence of if - else statements.
 * It runs the first case whose value is equal to the condition expression.
 *
 * Go's switch is like the one in C, C++, Java, JavaScript, and PHP,
 * except that Go only runs the selected case, not all the cases that follow.
 * In effect, the break statement that is needed at the end of each case in those languages
 * is provided automatically in Go. Another important difference is that Go's switch cases
 * need not be constants, and the values involved need not be integers.
 */

func main() {
	foo := 3

	// define classic switch
	switch foo {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case fooFunc():
		fmt.Println("Four")
	default:
		fmt.Println("None")
	}

	// do not use condition in switch
	// condition in case statements
	switch {
	case foo > 1:
		fmt.Println("One")
	case foo == 3:
		fmt.Println("Three")
		fallthrough
	case fooFunc() == 4:
		fmt.Println("Four")
	default:
		fmt.Println("None")
	}

	// e.g Not required break

	// *** Other Statements *** \\

	// Case list
	// switch c {
	// case ' ', '\t', '\n', '\f', '\r':
	// 		return true
	// }

	// Fallthrough - A fallthrough statement transfers control to the next case.
	// switch 2 {
	// case 1:
	// 		fmt.Println("1")
	// 		fallthrough
	// case 2:
	// 		fmt.Println("2")
	// 		fallthrough
	// case 3:
	// 		fmt.Println("3")
	// }
	// Output: 2 3

	// Exit with break
	// Loop:
	//   for _, ch := range "a b\nc" {
	//       switch ch {
	//       case ' ': // skip space
	//           break
	//       case '\n': // break at newline
	//           break Loop
	//       default:
	//           fmt.Printf("%c\n", ch)
	//       }
	//   }
	// Output: a b

	// Execution order
	// 	func Foo(n int) int {
	//     fmt.Println(n)
	//     return n
	//  }
	//
	// switch Foo(2) {
	// case Foo(1), Foo(2), Foo(3):
	// 		fmt.Println("First case")
	// 		fallthrough
	// case Foo(4):
	// 		fmt.Println("Second case")
	// }
	// Output:2 1 2 First case Second case
}

func fooFunc() int {
	return 4
}
