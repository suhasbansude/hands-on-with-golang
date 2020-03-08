package main

import (
	"fmt"
)

func main() {
	mySlice := []int{}
	// Initial Length and Capacity
	// Capacity doubled on append
	fmt.Printf("Capacity : %d\tLength : %d\n", cap(mySlice), len(mySlice))
	mySlice = append(mySlice, 1)
	fmt.Printf("Capacity : %d\tLength : %d\n", cap(mySlice), len(mySlice))
	mySlice = append(mySlice, 15)
	fmt.Printf("Capacity : %d\tLength : %d\n", cap(mySlice), len(mySlice))
	mySlice = append(mySlice, 25)
	fmt.Printf("Capacity : %d\tLength : %d\n", cap(mySlice), len(mySlice))
	mySlice = append(mySlice, 35)
	fmt.Printf("Capacity : %d\tLength : %d\n", cap(mySlice), len(mySlice))

	fmt.Println(mySlice)
}
