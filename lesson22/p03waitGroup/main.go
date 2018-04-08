package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	timeStart := time.Now()

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	timeEnd := time.Now()
	fmt.Println(timeEnd.Nanosecond() - timeStart.Nanosecond())
}

func foo() {

	for i := 0; i < 1000; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {

	for i := 0; i < 1000; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()
}
