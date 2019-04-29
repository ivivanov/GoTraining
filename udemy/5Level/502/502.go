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

	norman := person{
		name:      "norman",
		age:       10,
		icecreams: []string{"Banana", "Strawberry"}}

	// var people map[string]person = make(map[string]person)
	people := make(map[string]person, 1)
	people["ivan"] = ivan
	people["norman"] = norman

	// fmt.Println(people)
	for _, v := range people {
		fmt.Printf("\t%v loves:", v.name)
		for _, ice := range v.icecreams {
			fmt.Printf("%v,", ice)
		}
	}
}
