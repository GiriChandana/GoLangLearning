package main

import "fmt"

func Arrays() {
	var grades [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	fruits := [2]string{"apple", "jackfruit"}
	fmt.Println(fruits)

	for index, element := range fruits {
		fmt.Println(index, "=>", element)
	}

	//multidimensional arrays
	arr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[i][j])
		}
	}
}
