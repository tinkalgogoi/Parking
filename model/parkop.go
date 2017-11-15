package model

import (
	"fmt"
	"strconv"
)

var Cars = []Car{}
var Current_empty_slot int = 1
var Total_slots int
var err error
var leave_count int
var is_Slot_Empty bool = false

func StartParkingCar(args []string) {
	switch args[0] {
	case "create_parking_lot":
		Total_slots, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
		}
	case "park":
		if len(Cars) == Total_slots {
			fmt.Println("Sorry, parking lot is full")
		} else {
			i := 1
			if is_Slot_Empty {

				for i <= Total_slots {
					if i == Total_slots {
						c := Car{
							Slot_No: i,
							Car_No:  args[1],
							Colour:  args[2],
						}
						Cars = append(Cars, c)
						fmt.Println("Allocated slot number:", i)
						break
					}
					if i != Cars[i-1].Slot_No {
						c := Car{
							Slot_No: i,
							Car_No:  args[1],
							Colour:  args[2],
						}
						Cars = append(Cars[:i-1], append([]Car{c}, Cars[i-1:]...)...)
						fmt.Println("Allocated slot number:", i)
						if leave_count > 0 {
							leave_count--
							is_Slot_Empty = true
						}
						if leave_count <= 0 {
							is_Slot_Empty = false
						}
						break
					}
					i++
				}

			} else {
				c := Car{
					Slot_No: Current_empty_slot,
					Car_No:  args[1],
					Colour:  args[2],
				}
				Cars = append(Cars, c)
				fmt.Println("Allocated slot number:", Current_empty_slot)
				Current_empty_slot++
			}
		}
	case "leave":
		i, _ := strconv.Atoi(args[1])
		j := 0
		for _, c := range Cars {
			if c.Slot_No == i {
				Cars = append(Cars[:j], Cars[j+1:]...)
				is_Slot_Empty = true
				leave_count++
				fmt.Println("Slot No ", i, " is free")
			}
			j++
		}

	case "status":
		fmt.Println("Slot No.\t Registration No \t Colour")
		for _, c := range Cars {
			fmt.Println(c.Slot_No, "\t\t", c.Car_No, "\t\t", c.Colour)
		}
	case "registration_numbers_for_cars_with_colour":
		for _, c := range Cars {
			if c.Colour == args[1] {
				fmt.Printf("%s\t", c.Car_No)
			}
		}
		fmt.Printf("\n")
	case "slot_numbers_for_cars_with_colour":
		for _, c := range Cars {
			if c.Colour == args[1] {
				fmt.Printf("%s\t", strconv.Itoa(c.Slot_No))
			}
		}
		fmt.Printf("\n")
	case "slot_number_for_registration_number":
		for _, c := range Cars {
			if c.Car_No == args[1] {
				fmt.Println(c.Slot_No)
			}
		}
	default:
		fmt.Println("Incorrect Choice")
	}
}
