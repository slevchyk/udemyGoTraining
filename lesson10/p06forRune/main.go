package main

import "fmt"

func main() {

	for i := 32; i < 150; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	a := 'a'
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", rune(a))
}
