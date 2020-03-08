package main

import "fmt"

func main() {

	// Range with pointer semantic
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	fmt.Printf("Bfr[%s] : ", friends[1])

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1])
		}
	}

	// Range with value semantic
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", v)
		}
	}

	// While ranging over array range make saperate copy of array and iterate.
}
