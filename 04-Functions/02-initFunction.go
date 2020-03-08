package main

import "fmt"

func main() {
	fmt.Println("I am main function")
	fmt.Println(sayHello())
}

func sayHello() string {
	fmt.Println("I am sayHello Function")
	return "Hello"
}

func init() {
	fmt.Println("I am init Function")
}
