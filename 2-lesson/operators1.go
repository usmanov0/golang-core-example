package main

import "fmt"

func operators1() {
	a := 64
	b := 20

	// addition
	result1 := a + b
	fmt.Printf("a + b result of = %d", result1)

	// subtraction
	result2 := a - b
	fmt.Printf("\na - b result of = %d", result2)

	// multiplication
	result3 := a * b
	fmt.Printf("\na * b result of = %d", result3)

	// divide
	result4 := a / b
	fmt.Printf("\na / b result of = %d", result4)

	// modul
	result5 := a % b
	fmt.Printf("\na %% b result of = %d", result5)
}
