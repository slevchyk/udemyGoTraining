package main

import "fmt"

func main() {

	name := "Sideney"
	value, ok := name.(string)

	if ok {
		fmt.Printf("%q\n", value)
	} else {
		fmt.Println("it's not a string")
	}
}
