package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Circle struct {
	radius float64
}

// Implements Shape interface
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

// Implements Shape interface
func (r Rectangle) getArea() float64 {
	return r.height * r.width
}

func printArea(a Shape) {
	fmt.Println(a.getArea())
}

func main() {
	c := Circle{5}
	r := Rectangle{7, 8}
	// Print Circle Area
	printArea(c)
	// print Rectangle Area
	printArea(r)
}
