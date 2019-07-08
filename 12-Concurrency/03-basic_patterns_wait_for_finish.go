package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForFinished()
}

func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Microsecond)
		close(ch)
		fmt.Println("employee : send signal")
	}()

	_, wd := <-ch
	fmt.Println("manager : recv'd signal : ", wd)

	time.Sleep(time.Second)
	fmt.Println("-------------")
}
