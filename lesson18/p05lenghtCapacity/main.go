package main

import "fmt"

func main() {

	records := make([]int, 0, 5)

	fmt.Println("---------------------")
	fmt.Println("Len: ", len(records))
	fmt.Println("Cap: ", cap(records))
	fmt.Println("---------------------")

	for i := 0; i < 80; i++ {
		records = append(records, i)
		fmt.Println("Len: ", len(records), " Cap: ", cap(records), " Val: ", records[i])
	}
}
