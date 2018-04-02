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

func specs(anyType interface{}) {
	fmt.Println(anyType)
}

func main() {

	jack := dog{animal{"Woof"}, true}
	miclky := cat{animal{"Meow"}, false}

	specs(jack)
	specs(miclky)
}
