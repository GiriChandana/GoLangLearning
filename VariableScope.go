package main

import "fmt"

func VariableScope() {
	country := "India"
	{
		city := "Hyderabad"
		fmt.Printf("Country = %v, City = %v ", country, city)
		fmt.Println()
	}
	//fmt.Printf("Country = %v, City = %v ", country, city)
}
