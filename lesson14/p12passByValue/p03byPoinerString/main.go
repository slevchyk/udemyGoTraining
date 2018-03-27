package main

import "fmt"

func main() {

	var age string
	age = "Serhii"

	fmt.Println(age)
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(age)
	fmt.Println(&age)
}

func changeMe(x *string) {

	fmt.Println(*x)
	fmt.Println(x)

	*x = "Bill"

	fmt.Println(*x)
	fmt.Println(x)
}
