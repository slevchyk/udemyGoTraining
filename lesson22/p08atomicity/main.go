package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

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
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		atomic.AddInt64(&counter, 1)
		fmt.Println(name, ": ", i, " - ", counter)

	}
	wg.Done()
}
