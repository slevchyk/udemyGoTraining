package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {

	var p1 Person

	rdr := strings.NewReader(`{"First":"Serhii", "Last":"Levchyk", "Age":31}`)

	myDecoder := json.NewDecoder(rdr)
	myDecoder.Decode(&p1)

	fmt.Println(p1)
}
