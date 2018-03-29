package main

import "fmt"

func main() {

	word := "Go test"
	hash := hashText(word, 12)

	fmt.Println(hash)

}

func hashText(word string, bucket int) int {

	number := int(word[0])
	hash := number % bucket

	return hash
}
