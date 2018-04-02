package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	var p1 Person
	fmt.Println(p1)

	byteSlice := []byte(`{"First":"Serhii", "Last":"Levchyk", "Age":31}`)

	json.Unmarshal(byteSlice, &p1)

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
}
