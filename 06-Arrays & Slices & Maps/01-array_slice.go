package main

import "fmt"

func main() {
	// Array
	arr1 := [5]int{1, 2, 3, 4, 5}

	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr1)

	// Slice
	slice1 := []int{1, 2, 3, 4, 5}
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slice1)
	// Slices have length and capacity
	fmt.Println("Capacity : ", cap(slice1))
	fmt.Println("Length : ", len(slice1))
}
