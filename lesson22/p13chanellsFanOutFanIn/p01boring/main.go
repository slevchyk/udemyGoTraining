package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You are both boring; I'm leaving")

}

func boring(msg string) <-chan string {

	out := make(chan string)

	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return out
}

func fanIn(input1, input2 <-chan string) <-chan string {

	out := make(chan string)

	go func() {
		for {
			out <- <-input1
		}
	}()

	go func() {
		for {
			out <- <-input2
		}
	}()

	return out
}
