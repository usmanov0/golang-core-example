package main

import "fmt"

func operators4() {
	a := 14
	b := 30

	// & (bitwise AND)
	result1 := a & b
	fmt.Printf("a & b result of %d", result1)

	// | (bitwise OR)
	result2 := a | b
	fmt.Printf("\np | q result of %d", result2)

	// ^ (bitwise XOR)
	result3 := a ^ b
	fmt.Printf("\np ^ q result of %d", result3)

	// << (left shift)
	result4 := a << 1
	fmt.Printf("\np << 1 result of %d", result4)

	// >> (right shift)
	result5 := a >> 1
	fmt.Printf("\np >> 1 result of %d", result5)

	// &^ (AND NOT)
	result6 := a &^ b
	fmt.Printf("\np &^ q result of %d", result6)
}
