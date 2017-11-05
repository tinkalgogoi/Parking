package model

import (
	"fmt"
	"strconv"
)

var cars = []Car{}
var current_empty_slot int = 1
var total_slots int
var err error
var leave_count int
var Is_Slot_Empty bool = false

func StartParkingCar(args []string) {
	//fmt.Println("args : ", args)
	switch args[0] {
	case "create_parking_lot":
		arg1 := args[1]
		total_slots, err = strconv.Atoi(arg1)
		if err != nil {
			fmt.Println(err)
		}
	case "park":
		if len(cars) == total_slots {
			fmt.Println("Sorry, parking lot is full")
		} else {
			i := 1
			if Is_Slot_Empty {

				for i <= total_slots {
					if i != cars[i-1].Slot_No {
						c := Car{
							Slot_No: i,
							Car_No:  args[1],
							Colour:  args[2],
						}
						cars = append(cars[:i-1], append([]Car{c}, cars[i-1:]...)...)
						fmt.Println("Allocated slot number:", i)
						if leave_count > 0 {
							leave_count--
							Is_Slot_Empty = true
						}
						if leave_count <= 0 {
							Is_Slot_Empty = false
						}
						break
					}
					i++
				}

			} else {
				c := Car{
					Slot_No: current_empty_slot,
					Car_No:  args[1],
					Colour:  args[2],
				}
				cars = append(cars, c)
				fmt.Println("Allocated slot number:", current_empty_slot)
				current_empty_slot++
			}
		}
	case "leave":
		i, _ := strconv.Atoi(args[1])
		j := 0
		for _, c := range cars {
			if c.Slot_No == i {
				cars = append(cars[:j], cars[j+1:]...)
				Is_Slot_Empty = true
				leave_count++
				fmt.Println("Slot No ", i, " is free")
			}
			j++
		}

	case "status":
		fmt.Println("Slot No.\t Registration No \t Colour")
		for _, c := range cars {
			fmt.Println(c.Slot_No, "\t\t", c.Car_No, "\t\t", c.Colour)
		}
	case "registration_numbers_for_cars_with_colour":
		for _, c := range cars {
			if c.Colour == args[1] {
				fmt.Printf("%s\t", c.Car_No)
			}
		}
		fmt.Printf("\n")
	case "slot_numbers_for_cars_with_colour":
		for _, c := range cars {
			if c.Colour == args[1] {
				fmt.Printf("%s\t", strconv.Itoa(c.Slot_No))
			}
		}
		fmt.Printf("\n")
	case "slot_number_for_registration_number":
		for _, c := range cars {
			if c.Car_No == args[1] {
				fmt.Println(c.Slot_No)
			}
		}
	default:
		fmt.Println("Incorrect Choice")
	}
}
