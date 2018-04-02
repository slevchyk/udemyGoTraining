package main

import "fmt"

func main() {

	var age int
	age = 31

	fmt.Println(age)
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(age)
	fmt.Println(&age)
}

func changeMe(x *int) {

	fmt.Println(*x)
	fmt.Println(x)

	*x++

	fmt.Println(*x)
	fmt.Println(x)
}
