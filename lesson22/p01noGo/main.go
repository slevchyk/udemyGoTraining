package main

import (
	"fmt"
	"time"
)

func main() {

	timeStart := time.Now()

	foo()
	bar()

	timeEnd := time.Now()
	fmt.Println(timeEnd.Nanosecond() - timeStart.Nanosecond())
}

func foo() {

	for i := 0; i < 1000; i++ {
		fmt.Println("foo: ", i)
	}
}

func bar() {

	for i := 0; i < 1000; i++ {
		fmt.Println("bar: ", i)
	}
}
