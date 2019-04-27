package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	fmt.Printf("%v %v %v\n", x, y, z)
	fmt.Printf("%T %T %T\n", x, y, z)
}
