package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int `json:"MyAge"`
}

func main() {

	var p1 Person
	fmt.Println(p1)
	
	byteSlice := []byte(`{"First":"Serhii", "Last":"Levchyk", "MyAge":31}`)

	json.Unmarshal(byteSlice, &p1)

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
}

