package main

import (
	"fmt"
	"sort"
)

func main() {

	sort0()
	sort1()
	sort2()
	sort3()
}

func sort0() {

	intSlice := []int{1, 4, 9, 34, 0, 12, 7, 98, 54}
	fmt.Println(intSlice)

	sort.IntSlice(intSlice).Sort()
	fmt.Println(intSlice, "\n")

}

func sort1() {

	intSlice := []int{1, 4, 9, 34, 0, 12, 7, 98, 54}
	fmt.Println(intSlice)

	sort.Ints(intSlice)
	fmt.Println(intSlice, "\n")
}

func sort2() {

	intSlice := []int{1, 4, 9, 34, 0, 12, 7, 98, 54}
	fmt.Println(intSlice)

	sort.Sort(sort.IntSlice(intSlice))
	fmt.Println(intSlice, "\n")
}

func sort3() {

	intSlice := []int{1, 4, 9, 34, 0, 12, 7, 98, 54}
	fmt.Println(intSlice)

	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice, "\n")
}
