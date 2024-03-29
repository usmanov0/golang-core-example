// Select statements

package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Fast"
			time.Sleep(time.Millisecond * 10)
		}
	}()

	go func() {
		for {
			channel2 <- "Slow"
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)

		case message2 := <-channel2:
			fmt.Println(message2)

		default:
			fmt.Println("no data")
			time.Sleep(time.Millisecond * 300)
		}
	}
}
