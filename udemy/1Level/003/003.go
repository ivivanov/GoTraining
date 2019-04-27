package main

import "fmt"

var (
	x int    = 42
	y string = "James Bond"
	z bool   = true
)

func main() {
	s := fmt.Sprintf("%v %v %v\n", x, y, z)
	fmt.Println(s)
}
