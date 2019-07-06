package main

import (
	"fmt"
)

func main() {
	// Unbuffered channel
	ch := make(chan int)
	go calculateSum(ch, 10)
	fmt.Println(<-ch)
}

func calculateSum(ch chan int, num int) {
	sum := 0
	for i := 1; i < num; i++ {
		sum += i
	}
	ch <- sum
}
