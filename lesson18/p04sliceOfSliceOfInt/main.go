package main

import "fmt"

func main() {

	records := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, i+j)
		}
		records = append(records, transaction)
	}

	fmt.Println(records)
}
