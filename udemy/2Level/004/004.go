package main

import "fmt"

const number = 11

func main() {
	print(number)
	a := number << 1
	print(a)
}

func print(num int) {
	fmt.Printf("%06b\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%X\n", num)
}
