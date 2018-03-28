package main

import "fmt"

func main()  {

	records := make([][]string, 0)

	record01 := make([]string, 4)
	record01[0] = "Serhii"
	record01[1] = "Levchyk"
	record01[2] = "012"
	record01[3] = "100"

	records = append(records, record01)

	record02 := make([]string, 4)
	record02[0] = "Yulii"
	record02[1] = "Levchyk"
	record02[2] = "008"
	record02[3] = "069"

	records = append(records, record02)

	fmt.Println(records)

}
