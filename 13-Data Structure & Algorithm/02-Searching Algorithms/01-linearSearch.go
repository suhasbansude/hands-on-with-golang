package main

import (
	"fmt"
)

func main() {
	testArr := []int{1, 3, 6, 7, 9, 34, 78, 99, 65}
	fmt.Println(linearSearch(testArr, 99))
	fmt.Println(linearSearch(testArr, 900))
}

func linearSearch(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}
