package main

import "fmt"

func operators() {
	var a, b int = 10, 20

	fmt.Printf("a: %d, b: %d\n", a, b)
	fmt.Printf("a == b ? %t\n", a == b)
	fmt.Printf("a != b ? %t\n", a != b)

	fmt.Printf("a > b ? %t\n", a > b)
	fmt.Printf("a < b ? %t\n", a < b)

	fmt.Printf("a >= b ? %t\n", a >= b)
	fmt.Printf("a <= b ? %t\n", a <= b)

	fmt.Printf("!(a == b) = %t\n", !(a == b))

	var c, d bool = true, false
	fmt.Printf("c && d ? %t\n", (c && d))
	fmt.Printf("c || d ? %t\n", (c || d))

}
