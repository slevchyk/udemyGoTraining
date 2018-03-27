package main

import "fmt"

func main() {

	fmt.Println(sprintName("Serhii", "Levchyk"))
}

func sprintName(fName, lName string) (string, string) {

	return fmt.Sprint(fName, lName), fmt.Sprint(lName, fName)
}
