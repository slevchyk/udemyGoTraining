package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

type killer struct {
	person
	first         string
	licenseToKill bool
}

func main() {

	p1 := killer{
		person{
			"Serhii",
			"Levchyk",
			31},
		"Coud kill somebody",
		true}

	p2 := killer{
		person{
			"Yuliia",
			"Levchyk",
			27},
		"Coudn't kill nobody",
		false}

	fmt.Println(p1.first, p1.second, p1.licenseToKill)
	fmt.Println(p2.first, p2.second, p2.licenseToKill)
}
