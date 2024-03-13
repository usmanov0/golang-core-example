package main

import "fmt"

func main() {

	statuses := make(map[string]int)

	// add value to map:
	statuses["pending"] = 0
	statuses["inited"] = 1
	statuses["running"] = 2
	statuses["timedout"] = 3
	statuses["successful"] = 4
	statuses["failed"] = 5

	// retrieve from map:
	var successfulStatus = statuses["successful"]
	fmt.Println(successfulStatus)

	// delete element from map:
	delete(statuses, "timedout")
}
