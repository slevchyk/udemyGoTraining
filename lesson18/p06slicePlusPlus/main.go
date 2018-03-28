package main

import "fmt"

func main() {

	records := make([]int, 1)

	fmt.Println(records)
	records[0] = 7
	fmt.Println(records)
	records[0]++
	fmt.Println(records)
}
