package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	insert(arr, 3, 100)
	fmt.Println(arr)
}

/**
*	Insert value at position k
 */
func insert(arr []int, k int, val int) bool {
	if arr == nil || k > len(arr) {
		return false
	}

	for i := k - 1; i < len(arr)-1; i++ {
		temp := arr[i+1]
		arr[i+1] = temp
	}
	arr[k-1] = val
	return true
}
