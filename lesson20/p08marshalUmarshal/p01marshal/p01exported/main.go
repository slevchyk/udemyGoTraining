package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {

	p1 := Person{"Serhii", "Levchyk", 31, 0}
	byteSlice, _ := json.Marshal(p1)

	fmt.Println(byteSlice)
	fmt.Printf("%T\n", byteSlice)
	fmt.Println(string(byteSlice))
}
