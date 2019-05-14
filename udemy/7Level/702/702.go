package main

import "fmt"

type person struct {
	name string
	age  int
}

// receiver is of type pointer we can modify person
func (p *person) setName(name string) {
	p.name = name
}

// non pointer receiver - function can not modify person's fields
func (p person) getAge() int {
	p.age = 13 // we cant modify person. Those changes remain in the scope of the function
	return p.age
}

func main() {
	p := person{name: "Johnny", age: 12}
	fmt.Println(p)
	p.setName("Smith")
	fmt.Println(p.name)
	fmt.Println(p, "age:", p.getAge())
}
