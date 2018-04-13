package main

import "fmt"

func main() {

	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fmt.Println("factorial of ", n, " is ", <-factorial(n))

}

func factorial(n int) <-chan int {

	out := make(chan int)

	go func() {
		factorial := 1
		for i := n; i > 0; i-- {
			factorial *= i
		}
		out <- factorial
	}()

	return out
}
