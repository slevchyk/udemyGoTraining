package main

import (
	"sort"
	"fmt"
)

func main() {

	sort0()
	sort1()
	sort2()
	sort3()
}

func sort0() {

	stringSlice := []string{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Println(stringSlice)

	sort.StringSlice(stringSlice).Sort()
	fmt.Println(stringSlice, "\n")
}

func sort1() {

	stringSlice := []string{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Println(stringSlice)

	sort.Strings(stringSlice)
	fmt.Println(stringSlice, "\n")
}

func sort2() {

	stringSlice := []string{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Println(stringSlice)

	sort.Sort(sort.StringSlice(stringSlice))
	fmt.Println(stringSlice, "\n")
}

func sort3() {

	stringSlice := []string{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Println(stringSlice)

	sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
	fmt.Println(stringSlice, "\n")
}