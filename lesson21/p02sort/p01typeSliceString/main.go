package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {return len(p)}
func (p people) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p people) Less(i, j int) bool {return p[i] < p[j]}

func main() {

	sort0()
	sort1()
	sort2()
	sort3()
}

func sort0() {

	groupPeople := people{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Printf("%T\n", groupPeople)
	fmt.Println(groupPeople)

	sort.Sort(groupPeople)
	fmt.Println(groupPeople, "\n")
}

func sort1() {

	groupPeople := people{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Printf("%T\n", groupPeople)
	fmt.Println(groupPeople)

	sort.Strings([]string(groupPeople))
	fmt.Println(groupPeople, "\n")
}

func sort2() {

	groupPeople := people{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Printf("%T\n", groupPeople)
	fmt.Println(groupPeople)

	sort.Sort(sort.StringSlice([]string(groupPeople)))
	fmt.Println(groupPeople, "\n")
}

func sort3() {

	groupPeople := people{"Serhii", "Andrii", "Pavlo", "Oleh"}
	fmt.Printf("%T\n", groupPeople)
	fmt.Println(groupPeople)

	sort.Sort(sort.Reverse(sort.StringSlice([]string(groupPeople))))
	fmt.Println(groupPeople, "\n")
}