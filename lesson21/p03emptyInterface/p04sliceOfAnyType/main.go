package main

import "fmt"

type animal struct {
	Sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {

	jack := dog{animal{"Woof"}, true}
	miclky := cat{animal{"Meow"}, false}

	animals := []interface{}{jack, miclky}

	fmt.Println(animals)
}
