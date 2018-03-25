package main

import (
	"fmt"
	"github.com/slevchyk/udemyGoTraining/lesson02/stringutil"
)

func main() {

	fmt.Println("Origin name is",
		stringutil.Name,
		"and ",
		stringutil.Reverse(stringutil.Name),
		" is reversed \n",
	)
}
