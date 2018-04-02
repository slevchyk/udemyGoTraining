package main

import "fmt"

func main() {

	frKey := "fr"

	myGreeting := map[string]string{
		"en":  "Hello",
		"uk":  "Привіт",
		frKey: "Salut"}

	fmt.Println(myGreeting)

	if value, exis := myGreeting[frKey]; exis {
		delete(myGreeting, frKey)
		fmt.Println("Record ", frKey, " value ", value, "delted")
		fmt.Println(myGreeting)
	} else {
		fmt.Println("Can't find record with Key ", frKey)
	}

	if value, exis := myGreeting[frKey]; exis {
		delete(myGreeting, frKey)
		fmt.Println("Record ", frKey, " value ", value, "delted")
		fmt.Println(myGreeting)
	} else {
		fmt.Println("Can't find record with Key ", frKey)
	}

}
