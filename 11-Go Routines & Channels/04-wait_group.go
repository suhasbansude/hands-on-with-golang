package main

import (
	"fmt"
	"sync"
)

func printMessageHi() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hi")
	}
	wg.Done()
}

func printMessageHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go printMessageHi()
	go printMessageHello()
	wg.Wait()
}
