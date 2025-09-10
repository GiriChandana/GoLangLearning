package main

import (
	"fmt"
)

func Slices() {
	slice := []int{10, 20, 30, 40, 50}

	fmt.Println("slice: ", slice)
	fmt.Println("Length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))

	fmt.Println()

	sub_slice1 := slice[0:3]
	fmt.Println("sub_slice1: ", sub_slice1)
	fmt.Println("Length: ", len(sub_slice1))
	fmt.Println("Capacity: ", cap(sub_slice1))

	fmt.Println()

	sub_slice2 := slice[0:]
	fmt.Println("sub_slice2: ", sub_slice2)
	fmt.Println("Length: ", len(sub_slice2))
	fmt.Println("Capacity: ", cap(sub_slice2))

	fmt.Println()

	slice2 := make([]int, 3, 5)
	fmt.Println("slice2: ", slice2)
	fmt.Println("Length: ", len(slice2))
	fmt.Println("Capacity: ", cap(slice2))

	slice2_subslice1 := append(slice2, 12, 13, 13)
	fmt.Println(slice2_subslice1)
	fmt.Println("Length: ", len(slice2_subslice1))
	fmt.Println("Capacity: ", cap(slice2_subslice1))

	fmt.Println("Slice2 Length: ", len(slice2))
	fmt.Println("Slice2 Capacity: ", cap(slice2))

	//append one slice to another

	fmt.Println()

	arr := []int{1, 2, 3, 4, 5, 6}

	arr_slice1 := arr[:1]
	arr_slice2 := arr[4:]
	fmt.Println("arr_slice1: ", arr_slice1)
	fmt.Println("arr_slice2: ", arr_slice2)
	arr2 := append(arr_slice1, arr_slice2...)
	fmt.Println("append result of arr_slice1 and arr_slice2: : ", arr2)

	//copy slices
	arr3 := make([]int, 3)
	fmt.Println("arr3: ", arr3)
	copied := copy(arr3, arr2)
	fmt.Println("arr3: ", arr3)
	fmt.Println("arr2 copied to arr3: ", copied)

}
