package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

type killer struct {
	person
	licenseToKill bool
}

func main() {

	p1 := killer{
		person{
			"Serhii",
			"Levchyk",
			31},
		true}

	p2 := killer{
		person{
			"Yuliia",
			"Levchyk",
			27},
		false}

	fmt.Println(p1.licenseToKill)
	fmt.Println(p2.licenseToKill)
}
