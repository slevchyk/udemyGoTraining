package main

import "fmt"

func main() {

	var myGreeting= make(map[string]string)

	myGreeting["En"] = "Hello"
	myGreeting["Uk"] = "Привіт"

	fmt.Println(myGreeting)

}
