package main

import "fmt"

func main() {

	var didgit int32

	fmt.Print("Enter 'a' code didgit: ")
	fmt.Scanln(&didgit)

	if char := string(didgit); char == "a" {
		fmt.Println("Yes, it's 'a' code")
	}
}
