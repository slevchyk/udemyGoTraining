package main

import (
	"errors"
	"log"
	"math"
)

func main() {

	_, err := sqrt(-5)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}

	return math.Sqrt(f), nil
}
