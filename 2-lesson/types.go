// (basic types)

package main

import "fmt"

func types() {
	// integer
	var X uint8 = 225
	fmt.Println(X+1, X)
	var Y int16 = 32765
	fmt.Println(Y - 2)

	// float
	a := 20.45
	b := 34.89
	c := b - a
	fmt.Printf("Result: %f", c)
	fmt.Printf("\n c type of : %T", c)

	// complex
	var d complex128 = complex(6, 2)
	var e complex64 = complex(9, 2)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("d type of %T va e type of %T", d, e)

	// boolean
	str1 := "Str1"
	str2 := "str1"
	str3 := "str1"
	result1 := str1 == str2
	result2 := str1 == str3
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Printf("type of result1 %T va type of result2 %T", result1, result2)

	// string
	str := "Hello world"
	fmt.Printf("\nstr length of string %d", len(str))
	fmt.Printf("\nText: %s", str)
	fmt.Printf("\nstr type: %T", str)
	const PI = 3.14
}
