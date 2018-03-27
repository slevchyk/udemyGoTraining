package main

import "fmt"

func main() {

	var x [256]int

	fmt.Println(x)
	fmt.Println(len(x))

	for i := 0; i < 256; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("%v\t %T\t %b\n", v, v, v)

		if i > 50 {
			break
		}
	}
}
