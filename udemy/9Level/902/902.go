package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak(s string)
}

func (p *person) speak(s string) {
	fmt.Printf("name: %v, age: %v says: %v\n", p.name, p.age, s)
}

func main() {
	john := person{"jon", 22}
	saySomething(&john, "hello")
	// saySomething(john, "hello")\
}

func saySomething(h human, s string) {
	h.speak(s)
}
