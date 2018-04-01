package main

import (
	"fmt"
)

func main() {

	var val interface{} = 31

	fmt.Println(7 + val.(int))

}
