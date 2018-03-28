package main

import "fmt"

func main() {

	mySlice := []string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(mySlice)
	fmt.Println(mySlice[2:5])
	fmt.Println(mySlice[2])
	fmt.Println("mySlice"[2])

}
