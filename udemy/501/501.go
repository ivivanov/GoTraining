package main

import "fmt"

type person struct {
	name      string
	age       int
	icecreams []string
}

func main() {
	ivan := person{
		name:      "ivan",
		age:       15,
		icecreams: []string{"Vanilla", "Choco"}}

	for _, v := range ivan.icecreams {
		fmt.Println(v)
	}

	norman := person{
		name:      "norman",
		age:       10,
		icecreams: []string{"Banana", "Strawberry"}}

	fmt.Println(norman)
}
