package main

import (
	"fmt"
	"strconv"
)

func typeCasting() {
	var i int = 32
	var s string = strconv.Itoa(i)
	fmt.Printf("Var i:%v, Type:%T\n", i, i)
	fmt.Printf("Var s:%v, Type:%T\n", s, s)

	var b string = "56"
	a, err := strconv.Atoi(b)
	fmt.Printf("Var a:%v, Type:%T\n", a, a)
	fmt.Printf("Var err:%v, Type:%T\n", err, err)

}
