package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {

	p1 := Person{"Serhii", "Levchyk", 31, 1}

	myEncoder := json.NewEncoder(os.Stdout)
	err := myEncoder.Encode(p1)

	fmt.Println(err)
}
