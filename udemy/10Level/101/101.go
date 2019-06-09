package main

import (
	"fmt"
)

// solution a
// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 42
// 	}()

// 	fmt.Println(<-c)
// }

// solution b
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
