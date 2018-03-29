package main

import "fmt"

func main() {

	enKey := "en"
	ukKey := "uk"
	frKey := "fr"

	myGreeting := map[string]string{
		enKey: "Hello",
		ukKey: "Привіт",
		frKey: "Salut"}

	for key, val := range  myGreeting{
		fmt.Println("key: ", key, " - val:", val)
	}
}