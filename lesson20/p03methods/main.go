package main

import "fmt"

type person struct {
	first string
	second string
	age	int
}

func (p person) FullName() string {

	return p.first + " " + p.second
}

func main() {

	me := person{"Serhii", "Levchyk", 31}
	she := person{"Yuliia", "Levchyk", 27}

	fmt.Println(me.FullName(), me.age)
	fmt.Println(she.FullName(), she.age)
}