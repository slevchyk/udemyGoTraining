package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := increment(2)

	var total int

	for n := range c {
		total++
		fmt.Println(n)
	}

	fmt.Println("Total: ", total)

}

func increment(n int) chan string {

	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(j int) {

			for i := 0; i < 20; i++ {
				c <- fmt.Sprint("Process: ", j, " printing: ", i)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
			fmt.Println("Go ", i, " completed")
		}
		close(c)
	}()

	return (c)
}
