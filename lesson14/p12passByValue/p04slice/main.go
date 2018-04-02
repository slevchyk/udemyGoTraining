package main

import "fmt"

func main() {

	person := make([]string, 2, 25)

	person[0] = "Serhii"
	person[1] = "Bill"
	fmt.Println(person)

	changeMe(person)
	fmt.Println(person)
}

func changeMe(x []string) {

	x[0] = "Andrii"
}
