package main

import "fmt"

func main() {

	x := int(42)

	fmt.Println("x value is ", x)
	fmt.Println("x memmory address is ", &x)
	fmt.Printf("%d \n", &x)
}
