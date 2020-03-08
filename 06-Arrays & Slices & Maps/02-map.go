package main

import "fmt"

func main() {
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
