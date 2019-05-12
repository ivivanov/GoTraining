package main

import (
	"fmt"
	"math"
)

func main() {
	func(a int) {
		fmt.Println(a)
	}(2)

	cubic := func(number float64) float64 {
		return math.Pow(number, 3)
	}

	cubicOfTwo := cubic(2)
	fmt.Println(cubicOfTwo)

	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
