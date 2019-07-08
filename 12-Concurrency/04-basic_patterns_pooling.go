package main

import (
	"fmt"
	"time"
)

func main() {
	pooling()
}

func pooling() {
	ch := make(chan string)

	const emps int = 2

	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	const work int = 10

	for w := 0; w < work; w++ {
		ch <- "paper"
		fmt.Println("manager : send signal : ", w)
	}

	close(ch)
	fmt.Println("manager : send shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("------------")
}
