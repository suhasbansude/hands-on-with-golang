package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	c := getSum(a, b)
	fmt.Printf("Sum = %d", c)
}

func getSum(a, b int) int {
	return a + b
}
