package main

import (
	"fmt"
	"time"
)

func main() {
	// Drop word when buffer capacity is full
	drop()
}

func drop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : recv'd signal : ", p)
		}
	}()

	const work = 20

	for w := 0; w < 20; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : send signal : ", w)
		default:
			fmt.Println("manager : dropped data : ", w)
		}
	}

	close(ch)
	fmt.Println("manager : send shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------")
}
