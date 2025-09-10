package main

import "fmt"

func Pointers() {

	i := 10
	fmt.Printf("%T %v\n", &i, &i)
	fmt.Printf("%T %v\n", *(&i), *(&i))
}
