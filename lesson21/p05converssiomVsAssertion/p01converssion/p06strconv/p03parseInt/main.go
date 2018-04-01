package main

import (
	"strconv"
	"fmt"
)

func main() {

	b, _ := strconv.ParseBool("false")
	f, _ := strconv.ParseFloat("3.145", 64)
	i, _ := strconv.ParseInt("-42", 10, 32)
	u, _ := strconv.ParseUint("42", 10, 32)

	fmt.Println(b, f, i, u)

	p := strconv.FormatBool(true)
	r := strconv.FormatFloat(3.145, 'E', -1, 64)
	s := strconv.FormatInt(-42, 16)
	t := strconv.FormatUint(42, 16)

	fmt.Println(p, r, s, t)

}
