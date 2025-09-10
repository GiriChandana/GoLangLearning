package main

import (
	"fmt"
	"reflect"
)

func VariableType() {
	var a int = 5
	var b float32 = 10.5
	var c string = "Chandana"
	var d bool = true
	fmt.Printf("Variable a = %v is of type %T\n", a, a)
	fmt.Printf("Variable b = %v is of type %T\n", b, b)
	fmt.Printf("Variable c = %v is of type %T\n", c, c)
	fmt.Printf("Variable d = %v is of type %T\n", d, d)

	fmt.Printf("Type: %v\n", reflect.TypeOf("KodeCloud"))
	fmt.Printf("Type: %v\n", reflect.TypeOf(5))
	fmt.Printf("Type: %v\n", reflect.TypeOf(99.6))
	fmt.Printf("Type: %v\n", reflect.TypeOf(true))

}
