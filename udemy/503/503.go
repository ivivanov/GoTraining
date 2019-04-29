package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheels bool
}

func main() {
	sClass := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 5,
			color: "black"}}

	kamaz := truck{
		fourWheels: false,
		vehicle: vehicle{
			doors: 2,
			color: "white"}}

	fmt.Println(sClass)
	fmt.Println(kamaz)
}
