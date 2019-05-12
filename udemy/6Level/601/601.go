package main

import "fmt"

func main() {
	fmt.Println(foo())
	// fmt.Printf("%T\n", bar)
	v1, v2 := bar()
	fmt.Println(v1, v2)
}

func foo() int {
	return 32
}

func bar() (int, string) {
	return 42, "Wow!"
}
