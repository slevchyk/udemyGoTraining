package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first       string
	last        string
	age         int
}

func main() {

	p1 := Person{"Serhii", "Levchyk", 31}
	byteSlice, _ := json.Marshal(p1)

	fmt.Println(byteSlice)
	fmt.Printf("%T\n", byteSlice)
	fmt.Println(string(byteSlice))
}
