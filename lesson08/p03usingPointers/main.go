package main

import "fmt"

func main() {

	a := 31

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 32

	fmt.Println(a)
}
