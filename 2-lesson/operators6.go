// misc (boshqa) operatorlar

package main

import "fmt"

func operators6() {
	a := 4

	// use the address in memory (&) and pointer (*) operators
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
