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

	// Map
	map1 := make(map[string]string)
	fmt.Printf("%T\n", map1)

	// Add
	map1["name"] = "My First Map Example"

	// Read
	val, ok := map1["name"]
	if ok {
		fmt.Println(val)
	}

	// Delete
	delete(map1, "name")
	_, isVal := map1["name"]
	fmt.Println(isVal)
}
