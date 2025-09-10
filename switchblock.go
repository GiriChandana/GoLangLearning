package main

import "fmt"

func switchBlock() {

	var a int = 20

	switch a {
	case 10:
		fmt.Println(a)
	case 30, 20:
		fmt.Println("Either or case")
		fallthrough
	default:
		fmt.Println("default case")
	}

	day := "wednesday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
		fallthrough
	case "thursday":
		fmt.Println("thursday")
		fallthrough
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}
}
