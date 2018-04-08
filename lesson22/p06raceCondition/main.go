package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()

	fmt.Println("final couner: ", counter)

}

func increment(name string) {

	for i := 0; i < 45; i++ {

		x := counter
		x++

		time.Sleep(time.Duration(10 * time.Millisecond))
		counter = x
		fmt.Println(name, ": ", i, " - ", counter)

	}
	wg.Done()
}
