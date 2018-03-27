package main

import "fmt"

func main() {

	fmt.Println(sprintName("Serhii", "Levchyk"))
}

func sprintName(fName, lName string) (s string) {

	s = fmt.Sprint(fName, lName)
	return
}
