package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println(x, " - Int")
	case string:
		fmt.Println(x, " - String")
	case Contact:
		fmt.Println(x, " - Contact")
	default:
		fmt.Println(x, " - Unknown")
	}

}

func main() {

	SwitchOnType(7)
	SwitchOnType("Hello")

	var serhii Contact = Contact{"Wassup", "Bro"}
	SwitchOnType(serhii)

}
