package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

func info(x Shape) {
	fmt.Printf("Type is %T\n", x)
	fmt.Println(x)
	fmt.Println("Area is", x.area())
}

func totalArea(shapes ...Shape) float64 {

	var totalArea float64

	for _, val := range shapes {
		totalArea += val.area()
	}

	return totalArea
}

func main() {

	s := Square{10}
	c := Circle{5}

	info(s)
	info(c)

	fmt.Println("Total area is", totalArea(s, c))
}
