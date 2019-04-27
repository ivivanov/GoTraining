package main

import "fmt"

const (
	current  = 2019
	next     = current + iota
	nextNext = current + iota
)

func main() {
	fmt.Println(current)
	fmt.Println(next)
	fmt.Println(nextNext)
}
