package main

import "fmt"

func main() {
	testArr := []int{3, 5, 6, 7, 8, 9, 12, 66, 99}
	fmt.Println(binarySearch(testArr, 12))
	fmt.Println(binarySearch(testArr, 3))
	fmt.Println(binarySearch(testArr, 101))
}

func binarySearch(arr []int, num int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == num {
			return mid
		} else if arr[mid] < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
