package main

import "fmt"

func main() {

	fmt.Println(sprintName("Serhii", "Levchyk"))
}

func sprintName(fName, lName string) string{

	return fmt.Sprint(fName, lName)
}
