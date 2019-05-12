package main

import "fmt"

func main() {
	defer fmt.Println("1st line")
	bar()
	foo()
	fmt.Println("last line")

	/* output
	bar
	defer in bar
	foo
	last line
	1st line
	*/
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	defer func() {
		fmt.Println("defer in bar")
	}()

	fmt.Println("bar")
}
