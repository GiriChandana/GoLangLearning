package main

import "fmt"

var course string = "Kodecloud"

func VariableScope() {
	country := "India"
	{
		city := "Hyderabad"
		fmt.Printf("Country = %v, City = %v ", country, city)
		fmt.Println()
		fmt.Printf("Global Variable inside block = %v", course)
		fmt.Println()
	}
	//fmt.Printf("Country = %v, City = %v ", country, city)
	fmt.Printf("Global Variable outside block = %v", course)
}
