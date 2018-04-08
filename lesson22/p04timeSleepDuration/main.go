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

	for i := 0; i < 45; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {

	for i := 0; i < 45; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(30 * time.Millisecond))
	}
	wg.Done()
}
