package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)

	myGreeting["En"] = "Hello"
	myGreeting["Uk"] = "Привіт"

	fmt.Println(myGreeting)
}
