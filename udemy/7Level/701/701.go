package main

import "fmt"

func main() {
	foo := "foo"
	bar := "bar"

	fmt.Printf("[before] foo %v - %v\n", foo, &foo)
	fmt.Printf("[before] bar %v - %v\n", bar, &bar)

	// change value of foo from two different variables
	var tmp *string
	tmp = &foo
	fmt.Printf("%T with value %v\n", tmp, tmp)

	*tmp = "tmp"

	fmt.Printf("[after] foo %v - %v\n", foo, &foo)
}

func swapAddress(foo *string, bar *string) {
}
