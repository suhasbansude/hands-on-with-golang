package main

import (
	"fmt"
)

const c1 = 100

const (
	c2 = 200
	c3 = 300
)

const c4 int32 = 20

var (
	v1 = 0
)

var v2 = 10

func main() {
	// Shorthand to declare variable
	a := 15
	b, c := "Test", false

	fmt.Println(a, b, c)

	var d int = 15
	var e = 15
	var f, g float32 = -1, -2

	fmt.Println(d, e, f, g)

	fmt.Println("Constant")
	fmt.Println(c1, c2, c3, c4)

	fmt.Println(v1, v2)
}
