package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

func main() {

	me := person{"Serhii", "Levchyk", 31}
	she := person{"Yuliia", "Levchyk", 27}

	fmt.Println(me.first, me.second, me.age)
	fmt.Println(she.first, she.second, she.age)

}
