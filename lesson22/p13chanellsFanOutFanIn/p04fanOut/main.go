package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {

	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond)

}

// publisher pushes dadta to channel
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0

	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

// workerProcess reading data from channel
func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID

	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is %s\n", thisID, input)
	}
}
