package main

import "fmt"

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Shape interface {
	area() float64
}

func info(x Shape) {
	fmt.Printf("Type is %T\n", x)
	fmt.Println("Area is", x.area())
}

func main() {

	s := Square{10}

	info(s)

}
