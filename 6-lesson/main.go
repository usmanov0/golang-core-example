// for loop

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Method 1, a simple for loop:
	for i := 1; i < 8; i++ {
		fmt.Println(i)
	}

	// Method 2, infinite (unchecked) loop:
	for {
		fmt.Println("Alhamdulillah " + strconv.Itoa(time.Now().Second()))
		time.Sleep(1 * time.Second)
		if time.Now().Second() == 59 {
			break
		}

	}

	// 3-Method, while loop:
	j := 1
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// Method 4, iterating through collections using for range

	// Method 4, iterating through collections using for range
	myarr := [3]string{"yer", "quyosh", "oy"}
	for index, value := range myarr {
		fmt.Println("index:", index, "value:", value)
	}

	// loop through the elements of the map using the for range:
	mymap := map[int]string{
		1: "c#",
		2: "typescript",
		3: "go",
	}
	for key, value := range mymap {
		fmt.Println("key:", key, "value:", value)
	}
}
