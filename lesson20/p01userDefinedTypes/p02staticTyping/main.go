package main

import "fmt"

type foo int

func main() {

	var age foo
	age = 31
	fmt.Printf("%T %v \n", age, age)

	var yourAge int
	yourAge = 27
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// this doesn't work
	// fmt.Println(age + yourAge)

	fmt.Println(int(age) + yourAge)

}
