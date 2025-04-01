package main

import (
	"fmt"

	"github.com/TomasConesa/Go/interface/vehicles"
)

func main() {

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "BIKE"}

	var d float64
	for _, v := range vArray {
		fmt.Printf("Vehicle: %s \n", v)

		veh, err := vehicles.New(v, 400)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		distance := veh.Distance()
		fmt.Printf("Distance: %.2f \n", distance)
		fmt.Println()
		d += distance
	}
	fmt.Println()
	fmt.Printf("Total distance: %.2f", d)
}
