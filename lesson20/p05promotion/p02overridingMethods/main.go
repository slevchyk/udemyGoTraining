package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

func (p person) Greeting() {
	fmt.Println("I'm person")
}

type killer struct {
	person
	licenseToKill bool
}

func (k killer) Greeting() {
	fmt.Println("I'm killer")
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

	p1.Greeting()
	p2.person.Greeting()
}
