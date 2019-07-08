package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Microsecond)
		ch <- "paper"
		fmt.Println("employee : send signal")
	}()

	p := <-ch
	fmt.Println("manager : recv'd signal : ", p)

	time.Sleep(time.Second)
	fmt.Println("---------------")
}
