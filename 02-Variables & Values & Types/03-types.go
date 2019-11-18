package main

import "fmt"

type Mystring string

func main() {
	a := 10
	b := "movie"
	var c Mystring = "movie"
	var d int32 = 20

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	// 	Example - 1 - Compilation Error
	/* if b == c {
		fmt.Println("Both are equal")
	}*/

	// 	Example - 2 - Compilation Error
	/* if a == d {
		fmt.Println("Both are equal")
	}*/
}
