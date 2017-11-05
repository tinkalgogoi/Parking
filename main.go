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

	//n, err := fmt.Scanf("%d", &InputType)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(n)
	//}

	car := &model.Car{}

	vehicle := parking.ParkVehicle(car)
	InputType = 1
	vehicle.ParkOperation(InputType)

}
