package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {

	input := make(chan string)
	publisher(input)
	publisher(input)
	publisher(input)
	workerProcess(input)
	workerProcess(input)
	workerProcess(input)
	time.Sleep(1 * time.Millisecond)

}

func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0

	for {
		dataID++
		fmt.Printf("publisher %d is pushing dara\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID

	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is %s\n", thisID, input)
	}
}
