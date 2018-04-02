package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x = 12
	z := "I have a many: " + strconv.Itoa(x)

	fmt.Println(z)
}
