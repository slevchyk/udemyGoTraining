package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"En": "Hello",
		"Uk": "Привіт"}

	myGreeting["Fr"] = "Bonjour"

	fmt.Println(myGreeting)

	myGreeting["Fr"] = "Salut"
	fmt.Println(myGreeting)
}
