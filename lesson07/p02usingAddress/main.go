package main

import "fmt"

const coInchToCm float64 = 2.54

func main() {

	var inchLangth float64

	fmt.Print("Enter length in inch:")
	fmt.Scan(&inchLangth)

	fmt.Println(inchLangth, " inch in cm is ", inchLangth*coInchToCm)

}
