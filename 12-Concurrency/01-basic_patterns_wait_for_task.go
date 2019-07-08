package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForTask()
}

func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("Function Received Signal : ", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager : Send signal")

	time.Sleep(1)
	fmt.Println("End")
}
