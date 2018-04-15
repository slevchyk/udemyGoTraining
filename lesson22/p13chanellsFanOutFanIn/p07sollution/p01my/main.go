package main

import (
	"fmt"
	"sync"
)

func main() {

	chans := make([]<-chan int, 0)

	min := 0
	max := 1000

	c := gen(min, max)

	//FAN OUT
	//one channle couple go routines
	for i := min; i < max; i++ {
		if i%100 == 0 {
			cOut := factorial(c)
			chans = append(chans, cOut)
		}
	}

	//FAN IN
	//couple chans  merge to one chan
	f := merge(chans)
	for n := range f {
		fmt.Println(n)
	}
}

func gen(min, max int) <-chan int {
	out := make(chan int)
	go func() {
		for i := min; i < max; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(c []<-chan int) <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(c))

	for _, val := range c {
		go func(cIn <-chan int) {
			for n := range cIn {
				out <- n
			}
			wg.Done()
		}(val)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
