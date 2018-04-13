package main

import (
	"fmt"
)

func main() {

	out := make(chan uint64)

	go func() {
		for i := 0; i < 100; i++ {
			f := factorial(i)
			out <- uint64(<-f)
		}
		close(out)
	}()

	for n := range out {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {

	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}
