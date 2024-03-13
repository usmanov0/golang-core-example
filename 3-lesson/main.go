package main

import (
	"fmt"
	"time"
)

func main() {
	testSwitch2()
}

func testIf() {
	i := 7
	if i == 7 {
		fmt.Println("Seven!")
	}
}

func testIfElse() {
	points := 20000
	if points < 5000 {
		fmt.Println("Silver")
	} else if points < 10000 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Platinum")
	}
}

func testSwitch1() {
	weekday := time.Now().Weekday()
	switch weekday {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Mistake")
	}
}

func testSwitch2() {
	var greeting int
	_, err := fmt.Scan(&greeting)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	switch {
	case greeting == 1:
		fmt.Println("Assalomu alaykum")
	case greeting == 2:
		fmt.Println("أسلم عليكم")
	case greeting == 3:
		fmt.Println("Hello")
	case greeting == 4:
		fmt.Println("Bonjour")
	case greeting == 5:
		fmt.Println("Привет")
	default:
		fmt.Println("Xato")
	}
}

func testSwitch3() {
	var userChoice string = "uch"
	switch userChoice {
	case "bir":
		fmt.Println("C#")
	case "ikki", "uch":
		fmt.Println("Go")
	case "to'rt", "besh", "olti":
		fmt.Println("TypeScript")
	}
}
