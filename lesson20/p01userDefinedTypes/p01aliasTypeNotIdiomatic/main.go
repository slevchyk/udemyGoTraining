package main

import "fmt"

type foo int

func main() {

	var age foo
	age = 31

	fmt.Printf("%T %v \n", age, age)

}
