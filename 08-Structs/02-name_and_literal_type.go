package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
}

func main() {
	// literal type
	t1 := struct {
		Name string
		Age  int
	}{"John", 23}

	fmt.Println(t1)
	// name type
	var s Student
	var e Employee

	// Example1 - (Compilation Fail)
	/*
		s = e
		fmt.Println(s)
	*/

	// Example2 - (Explict type conversion required for name type)
	/*
		s = Student(e)
		fmt.Println(s)
	*/

	// Example3 - (Explict type conversion not required for literal type)
	s = t1
	e = t1
	fmt.Println(s)
	fmt.Println(e)

}
