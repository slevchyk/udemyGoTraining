package main

import "fmt"

type Square struct {
	side float64
}

func (x Square) area() float64 {
	return x.side * x.side
}

func main() {

	x := Square{10}
	fmt.Println(x.side)
	fmt.Println(x.area())

}
