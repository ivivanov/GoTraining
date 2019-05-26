package main

import (
	"fmt"
	"strings"
)

type person struct {
	first     string
	last      string
	age       int
	icecreams []string
}

// uses string builder
func (p *person) speak() {
	var str strings.Builder
	str.WriteString(fmt.Sprint(p.first, " ", p.last, " ", "age:", p.age, " "))

	for _, v := range p.icecreams {
		str.WriteString(fmt.Sprint(fmt.Sprint(v, " ")))
	}

	fmt.Println(str.String())
}

func (p *person) SetFirst(first string) {
	p.first = first
}

func (p *person) copy() person {
	// clone := person{first: p.first, last: p.last, age: p.age}
	clone := *p
	return clone
}

func (p *person) copy2() person {
	clone := person{first: p.first, last: p.last, age: p.age}
	copy(clone.icecreams, p.icecreams)
	return clone
}

func (p *person) AddIcecream(s string) {
	p.icecreams = append(p.icecreams, s)
}

func main() {
	ivan := person{first: "Ivan", last: "Vladimirov", age: 29, icecreams: []string{"Vanilla", "Choko", "Berry"}}
	/* 	Test 1
	ivan.speak()
	cloned := ivan.copy()
	ivan.AddIcecream("Fresh")
	ivan.speak()
	cloned.speak() */

	/* 	Test 2
		cloned := ivan.copy2()
	   	ivan.speak()
	   	ivan.AddIcecream("Fresh")
	   	ivan.speak()
	   	cloned.speak() */

	t1 := ivan.copy()

	fmt.Printf("t1 compared to ivan: %v, types - %T, %T\n", &t1 == &ivan, t1, ivan)
	fmt.Printf("icecreams compare - %v\n", &t1.icecreams == &ivan.icecreams)
}
