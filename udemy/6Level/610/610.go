package main

import "fmt"

func main() {
	a := foo(1)
	b := foo(1)
	c := foo(5)
	fmt.Print(a(), b(), b(), b(), "\n")
	fmt.Print(c(), c(), c())
}

func foo(step int) func() int {
	var i int
	return func() int {
		i += step
		return i
	}
}
