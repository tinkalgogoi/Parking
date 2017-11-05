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

			text = strings.Replace(text, "\n", "", -1)
			// for windows cmd
			//text = strings.Replace(text, "\r\n", "", -1)

			args := strings.Split(text, " ")
			StartParkingCar(args)
		}
	} else if command == 2 {
		fmt.Println("Enter the name of input file, Here its input.txt :")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		text = strings.Replace(text, "\n", "", -1)
		//for windows cmd
		//text := strings.Replace("input.txt", "\r\n", "", -1)

		if file, err := os.Open(text); err == nil {
			// make sure it gets closed
			defer file.Close()
			// create a new scanner and read the file line by line
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				//fmt.Println(scanner.Text())
				args := strings.Split(scanner.Text(), " ")
				//fmt.Println(args[0]);
				StartParkingCar(args)
			}
		}
	}
}
func (v *Van) ParkOperation(command int) {
	fmt.Println("Implement logic for van parking")
}
