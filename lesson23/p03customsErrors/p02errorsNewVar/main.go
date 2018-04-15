package main

import (
	"errors"
	"log"
	"math"
)

//ErrNorgatemath test error
var ErrNorgatemath = errors.New("norgate math: square root of negative number")

func main() {

	_, err := sqrt(-5)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgatemath
	}

	return math.Sqrt(f), nil
}
