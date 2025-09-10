package main

import "fmt"

func Structs() {
	type Student struct {
		name   string
		rollNo int
	}
	s := Student{
		name:   "Joe",
		rollNo: 23,
	}
	fmt.Printf("%+v", s)
}
