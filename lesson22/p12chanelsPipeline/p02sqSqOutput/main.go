package main

import (
	"fmt"
)

func main() {

	for n := range sq(sq(gen(3, 4))) {
		fmt.Println(n)
	}

}

func gen(n ...int) chan int {

	out := make(chan int)

	go func() {
		for _, val := range n {
			out <- val
		}
		close(out)
	}()

	return out
}

func sq(in chan int) chan int {

	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
