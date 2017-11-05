package model

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

func (c *Car) ParkOperation(cammand int) {

}
func (v *Van) ParkOperation(cammand int) {

}
