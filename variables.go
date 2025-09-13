package main

import "fmt"

func PrintVariables() {
	var a int = 5
	fmt.Println("int a = ", a)
	var b string = "Hello world using a variable"
	fmt.Println("string b = ", b)
	var c bool = false
	fmt.Println("boolean c =", c)
	var d float32 = 5
	fmt.Println("float32 d = ", d)
	var s, t string
	s = "Kodecloud platform"
	t = "courses"
	fmt.Printf("String s = %v, String t = %v", s, t)
	fmt.Println()
	x := 13.5
	fmt.Printf("flat x  = %f", x)

}
