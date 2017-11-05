package main

import (
	"Parking/model"
	"Parking/parking"
	"fmt"
)

var InputType int

func main() {
	fmt.Println("______________ MULTI-STOREY PARKING LOT MANAGEMENT SYSTEM ______________")
	fmt.Println("Input Type: 1 for command | 2 for file")

	_, err := fmt.Scanf("%d", &InputType)
	if err != nil {
		fmt.Println(err)
	}

	car := &model.Car{}

	vehicle := parking.ParkVehicle(car)
	vehicle.ParkOperation(InputType)

}
