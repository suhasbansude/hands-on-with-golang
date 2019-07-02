package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("Type-1")
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	fmt.Println("Type-2")
	j := 0
	for j < len(nums) {
		fmt.Println(nums[j])
		j++
	}

	fmt.Println("Type-3")
	for i, v := range nums {
		fmt.Println(i, v)
	}

	/*
		fmt.Println("Type-4")
		for {
			fmt.Println("I am Infinite loop")
		} */
}
