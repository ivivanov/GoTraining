package main

import "fmt"

func main() {
	fmt.Println(SayHi())
}

// SayHi prints Hi
func SayHi() string {
	return "Hi"
}

func notCoveredFunc() {
	fmt.Println(":(")
}
