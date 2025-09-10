package main

import "fmt"

func controlFlow() {
	var a string = "apple"

	if a == "app" {
		fmt.Println("found")
	} else if a == "appl" {
		fmt.Println("found")
	} else {
		fmt.Println("found")
	}

	var x, y = "kolkata", "kolkata"
	if x == y {
		fmt.Println("Strings are equal")
	} else if x != y {
		fmt.Println("Strings are not equal")
	}
	fmt.Println("thank you!")
}
