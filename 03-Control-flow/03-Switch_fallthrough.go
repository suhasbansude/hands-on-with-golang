package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// Switch Example
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// Switch with no condition
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")

	case t.Hour() < 17:
		fmt.Println("Good Afternoon")

	default:
		fmt.Println("Good Evening")
	}

	// Switch Example with Fallthrough
	test := 1

	switch test {
	case 1:
		fmt.Println("One")
		fallthrough
	case 0:
		fmt.Println("Zero")
	default:
		fmt.Printf("Not In O 1")
	}

}
