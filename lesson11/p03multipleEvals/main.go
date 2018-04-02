package main

import "fmt"

func main() {

	var digit int8

	fmt.Print("Enter didhit between 1 and 5: ")
	fmt.Scan(&digit)

	switch digit {
	case 1, 2:
		fmt.Println("You enter 1 or 2")
	case 3, 4:
		fmt.Println("You enter 3 or 4")
	case 5:
		fmt.Println("You enter 5")
	default:
		fmt.Println("False")
	}
}
