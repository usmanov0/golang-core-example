// Goroutines

package main

import (
	"fmt"
	"time"
)

func main() {
	go display("assalamu alaykum")
	go display("va alaykum assalam")

	fmt.Scanln()
}

func display(input string) {
	for i := 1; true; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
