package main

import "fmt"

// a function is said to be recursive if it calls itself.
func main() {
	helloWorld(10)
}

func helloWorld(num int) {
	if num == 0 {
		return
	}
	fmt.Println("Hello World : ", num)
	helloWorld(num - 1)
}
