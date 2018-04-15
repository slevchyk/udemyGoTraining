package main

import (
	"fmt"
	"log"
	"math"
)

//NorgateMathError new my own error type
type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("error occored: %v %v %v", n.lat, n.long, n.err)
}

func main() {

	_, err := sqrt(-5)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	return math.Sqrt(f), nil
}
