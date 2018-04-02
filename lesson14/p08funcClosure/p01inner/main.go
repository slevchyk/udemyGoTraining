package main

import "fmt"

func main() {

	x := 31
	fmt.Println(x)

	{
		fmt.Println(x)
		y := "only this block has access"
		fmt.Println(y)
	}

}
