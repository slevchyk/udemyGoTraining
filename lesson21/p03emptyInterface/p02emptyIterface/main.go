package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	Vehicle vehicle
	Doors   int
	Wheels  int
}

type plane struct {
	Vehicle vehicle
	Jet     bool
}

type boat struct {
	Vehicle vehicle
	Length  int
}

func main() {

	bmw := car{}
	audi := car{}
	vw := car{}

	boeing747 := plane{}
	airbus320 := plane{}
	an148 := plane{}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}

	rides := []vehicles{bmw, audi, vw, boeing747, airbus320, an148, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
