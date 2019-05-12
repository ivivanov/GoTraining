package main

import (
	"fmt"
	"math"
)

type square struct {
	a float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return math.Pow(s.a, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	sq := square{a: 15}
	info(sq)

	c := circle{r: 12.345}
	info(c)
}
