package main

import (
	"fmt"
)

func main() {

	var name interface{} = 31
	value, ok := name.(string)

	if ok {
		fmt.Printf("%q\t%T\n", value, value)
	} else {
		fmt.Println("it's not a string")
	}
}
