package main

import "fmt"

type vehicle struct {
	Seats int
	MaxSpeed int
	Color string
}

type car struct {
	Vehicle vehicle
	Doors int
	Wheels int
}

type plane struct {
	Vehicle vehicle
	Jet bool
}

type boat struct {
	Vehicle vehicle
	Length int
}

func main() {

	bmw := car{}
	audi := car{}
	vw := car{}
	cars :=[]car{bmw, audi, vw}

	boeing747 := plane{}
	airbus320 := plane{}
	an148 := plane{}
	planes := []plane{boeing747, airbus320, an148}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	boats := []boat{sanger, nautique, malibu}

	for key, value := range cars{
		fmt.Println(key, " - ", value)
	}

	for key, value := range planes{
		fmt.Println(key, " - ", value)
	}

	for key, value := range boats{
		fmt.Println(key, " - ", value)
	}

}
