package main

import "fmt"

func do(nums []int, callback func(n int))  {

	for _, n := range nums{
		callback(n)
	}
}

func main() {

	nums := []int{1, 2, 3, 5, 8, 13, 21}
	do(nums, func(n int) {

		fmt.Println(n)
	})

}