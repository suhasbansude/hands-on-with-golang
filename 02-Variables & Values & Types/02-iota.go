package main

import "fmt"

const (
	A1 = iota
	A2
	A3
)

const (
	B1 = iota
	B2
	B3
)

const (
	C1 = iota
	C2 = iota
	C3 = iota
)

const (
	D1 = iota
	D2
	_
	D3
)

const (
	E1 = iota + 2 // Strat at 2
	E2            // Increment by 1
	E3
)

const (
	P1 = iota + iota
	P2 = iota
)

func main() {
	fmt.Println(A1, A2, A3)
	fmt.Println(B1, B2, B3)
	fmt.Println(C1, C2, C3)
	// Skip value
	fmt.Println(D1, D2, D3)

	// Incremental Value
	fmt.Println(E1, E2, E3)
	fmt.Println(P1, P2)
}
