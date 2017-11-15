package model

import (
	"strconv"
	"testing"
)

var (
	createCmd = "create_parking_lot"
	createNo  = "6"
	parkCmd   = "park"
	car1No    = "KA-01-HH-1234"
	color1    = "White"
	leaveCmd  = "leave"
)

func TestCreateParkingLot(t *testing.T) {
	args := []string{createCmd, createNo}
	StartParkingCar(args)
	gotSlot, err := strconv.Atoi(args[1])
	if err != nil {
		t.Errorf("Got Error :", err)
	}
	if Total_slots != gotSlot {
		t.Errorf("Expected Slot no.s: %d and got: %d", Total_slots, gotSlot)
	}
}

func TestParking(t *testing.T) {
	args := []string{createCmd, createNo}
	StartParkingCar(args)

	args = []string{parkCmd, car1No, color1}
	StartParkingCar(args)
	StartParkingCar(args)
	StartParkingCar(args)
	StartParkingCar(args)
	if len(Cars) != 4 {
		t.Errorf("Expected parked cars: 4 and got: %d", len(Cars))
	}

}

func TestLeave(t *testing.T) {
	args := []string{createCmd, createNo}
	StartParkingCar(args)

	args = []string{parkCmd, car1No, color1}
	StartParkingCar(args)
	StartParkingCar(args)
	StartParkingCar(args)
	StartParkingCar(args)
	args = []string{leaveCmd, "2"}
	StartParkingCar(args)
	for _, v := range Cars {
		if v.Slot_No == 2 {
			t.Errorf("Expected car at slot: 2 to leave but the car still present: %s", v.Car_No)
		}
	}
}
