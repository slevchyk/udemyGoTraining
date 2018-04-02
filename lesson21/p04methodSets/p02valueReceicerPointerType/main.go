package main

import (
	"fmt"
	"math"
)

type circle struct {
	Radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func info(s Shape) {
	fmt.Println(s, s.Area())
}

func main() {

	c := circle{10}

	info(&c)
}
