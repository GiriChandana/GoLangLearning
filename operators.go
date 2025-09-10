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

}
