package main

import (
	"fmt"
	"sort"
)

func main() {
	testSlices3()
}

func testArrays1() {
	// named myarr and of 3 string types
	// create a new array of elements:
	var myarr [3]string

	// when assigning a value to an element, its index is written
	// and indexing starts at 0:
	myarr[0] = "GO"
	myarr[1] = "Programming"
	myarr[2] = "Hello"

	// the index is also used to get the value of the element:
	fmt.Println("Elements:")
	fmt.Println("1: ", myarr[0])
	fmt.Println("2: ", myarr[1])
	fmt.Println("3: ", myarr[2])
}

func testArrays2() {
	// short array declaration:
	myarr := [3]int{2, 4, 8}
	fmt.Println(myarr)
}

func testArrays3() {
	// compare rows
	myarr1 := [...]int{9, 7, 6, 5}
	myarr2 := [4]int{9, 7, 6, 5}
	fmt.Println(myarr1 == myarr2)
}

func testArrays4() {
	// 2d arrays
	myarr := [3][3]string{{"C#", "GO", "TypeScript"},
		{"Java", "C++", "Python"},
		{"Dart", "Kotlin", "Swift"}}
	fmt.Println(myarr[0][1])
}

func testArrays5() {
	myarr1 := [3]int{2, 4, 8}

	// to copy the entire string:
	myarr2 := myarr1
	fmt.Println(myarr1)
	fmt.Println(myarr2)

	myarr1[2] = 100
	fmt.Println(myarr1)
	fmt.Println(myarr2)
}

func testArrays6() {
	myarr1 := [3]int{2, 4, 8}

	// copy by reference from an array:
	myarr2 := &myarr1
	fmt.Println(myarr1)
	fmt.Println(*myarr2)

	myarr1[2] = 100
	fmt.Println(myarr1)
	fmt.Println(*myarr2)
}

func testSlices1() {
	// declare new slice:
	myslice := []int{2, 4, 8}

	// add a new element to the end of the slice:
	myslice = append(myslice, 10)

	// retirieve slice length
	fmt.Printf("len of slice: : %d", len(myslice))
}

func testSlices2() {
	myarr := [7]string{"olma", "anor", "shaftoli", "gilos",
		"o'rik", "uzum", "anjir"}

	// create a slice from an array:
	myslice := myarr[1:4]
	fmt.Println(myslice)
}

func testSlices3() {
	// sorting slice
	myslice := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}
	sort.Ints(myslice)
	fmt.Println(myslice)
}
