package main

import "fmt"

func VariadicFunction(name string, subject ...string) {
	fmt.Println("Hey ", name, "Your subjects are: ")
	for _, value := range subject {
		fmt.Print(value, " ")
	}
	fmt.Println()

}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
