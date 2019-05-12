package main

import "fmt"

func main() {
	foo("argument is anonym func which prints numbers from 1 to 10", func() {
		for i := 1; i < 11; i++ {
			fmt.Print(i, " ")
		}
	})

	foo("argument is ", bar)
}

func foo(s string, cb func()) {
	defer cb()
	fmt.Println(s)
}

func bar() {
	fmt.Println("bar")
}
