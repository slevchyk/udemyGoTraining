package main

import (
	"fmt"
)

func main() {

	rem := 12.31

	fmt.Printf("&T\n", rem)
	fmt.Printf("&T\n", int(rem))

	var val interface{} = 12

	fmt.Printf("&T\n", val)
	fmt.Printf("&T\n", val.(int))

}
