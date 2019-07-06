package main

import (
	"fmt"
	"time"
)

func main() {
	go printEvenNumbers(10)
	time.Sleep(10)
}

func printEvenNumbers(num int) {
	for i := 1; i < num; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
