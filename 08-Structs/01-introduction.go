package main

import (
	"fmt"
)

type Animal struct {
	Name  string
	Color string
	Age   float32
}

func main() {
	dog := Animal{"Tom", "Black", 4}

	fmt.Printf("Name %v\n", dog.Name)
	fmt.Printf("Color %v\n", dog.Color)
	fmt.Printf("Age %v\n", dog.Age)
	fmt.Println(dog)
}
