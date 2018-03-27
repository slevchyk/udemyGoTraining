package main

import "fmt"

func main() {
	
	person := make(map[string]int)

	person["Bill"] = 23
	person["Serhii"] = 31
	fmt.Println(person)

	changeMe(person)
	fmt.Println(person)
	
}

func changeMe(x map[string]int)  {

	x["Serhii"] = 32
	
}
