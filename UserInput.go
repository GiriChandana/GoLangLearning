package main

import "fmt"

func UserInput() {
	var name string
	var course string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s ", &name)
	fmt.Println("Enter course name: ")
	fmt.Scanf("%s", &course)
	fmt.Printf("Hey %v, you are enrolled into %v!", name, course)
}
