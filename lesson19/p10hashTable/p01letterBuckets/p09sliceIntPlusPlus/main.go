package main

import "fmt"

func main() {

	buckets := make([]int, 1)

	fmt.Println(buckets)
	buckets[0] = 7
	fmt.Println(buckets)
	buckets[0]++
	fmt.Println(buckets)
}
