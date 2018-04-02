package main

import "fmt"

func main() {

	avt := avarage(12, 45, 56, 178, 100)
	fmt.Println(avt)

}

func avarage(x ...float64) float64 {

	var total float64

	for _, v := range x {
		total += v
	}

	return total / float64(len(x))

}
