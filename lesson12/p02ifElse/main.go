package main

import "fmt"

func main() {

	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age < 10 {
		fmt.Println("Your age is less 10 years")
	} else if age >= 10 && age < 40 {
		fmt.Println("Your age is 10 or more but less 40")
	} else if age >= 40 && age < 60 {
		fmt.Println("Your age is 40 or more but less 60")
	} else {
		fmt.Println("Your age is 60 or more")
	}
}
