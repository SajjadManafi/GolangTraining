package main

import "fmt"

type Airplane struct {
	Vehicle
	VehicleType string
}

type Vehicle struct {
	name           string
	passengerCount int
}

func main() {

	v := Airplane{
		Vehicle: Vehicle{
			name:           "Boeing 737",
			passengerCount: 215,
		},
		VehicleType: "Air vehicle",
	}

	fmt.Println(v)

	fmt.Println(v.name, v.passengerCount, v.VehicleType)
	// or -> fmt.Println(v.Vehicle.name, v.Vehicle.passengerCount, v.VehicleType)

	// change name
	v.name = "Boeing"
	fmt.Println(v)

}
