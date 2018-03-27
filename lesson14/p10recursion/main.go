package main

import "fmt"

func factorial(x int) int {

	if x == 0 {
		return 1
	}

	return x * factorial(x - 1)
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Factorial of ", i, " is ", factorial(i))
	}
}
