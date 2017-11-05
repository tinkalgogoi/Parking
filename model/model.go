package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Car struct {
	Slot_No int
	Car_No  string
	Colour  string
}

type Van struct {
	Slot_No int
	Car_No  string
	Colour  string
}

func (c *Car) ParkOperation(command int) {
	if command == 1 {
		fmt.Println("Enter commands :")
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			// convert CRLF to LF for linux
			//text = strings.Replace(text, "\n", "", -1)
			// for windows
			text = strings.Replace(text, "\r\n", "", -1)
			//fmt.Println("text :", text)
			args := strings.Split(text, " ")
			//fmt.Println("args :", args)
			StartParkingCar(args)
		}
	} else if command == 2 {
		fmt.Println("Implement file input")
	}
}
func (v *Van) ParkOperation(command int) {
	fmt.Println("Implement logic for van parking")
}
