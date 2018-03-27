package main

import "fmt"

func filter(numbers []int, callback func(n int) bool) []int {

	var x []int

	for _, n := range numbers {
		if callback(n) {
			x = append(x, n)
		}
	}

	return x
}

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	x = filter(x, func(n int) bool{
		return n % 2 == 0
	})

	fmt.Println(x)
}