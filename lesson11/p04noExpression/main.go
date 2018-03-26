package  main

import "fmt"

func main() {

	var digit int8

	fmt.Print("Enter didhit between 1 and 5: ")
	fmt.Scan(&digit)

	switch {
	case digit == 1:
		fmt.Println("You enter 1")
	case digit == 2:
		fmt.Println("You enter 2")
	case digit == 3:
		fmt.Println("You enter 3")
	case digit == 4:
		fmt.Println("You enter 4")
	case digit == 5:
		fmt.Println("You enter 5")
	default:
		fmt.Println("False")
	}
}
