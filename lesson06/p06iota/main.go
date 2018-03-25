package main

import "fmt"

const (
	_ = iota
	A = iota * 10
	B = iota * 20
)

func main() {

	fmt.Println(A)
	fmt.Println(B)
}