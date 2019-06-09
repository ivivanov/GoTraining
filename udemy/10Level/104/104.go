package main

import (
	"fmt"
)

func main() {
	quit := make(chan bool)
	q := make(chan int)
	c := gen(q, quit)

	receive(c, q, quit)

	fmt.Println("about to exit")
}

func gen(q chan<- int, quit chan<- bool) <-chan int {
	c := make(chan int)

	go func() {
		q <- 222

		for i := 0; i < 100; i++ {
			c <- i
		}

		q <- 333
		close(quit)
	}()

	return c
}

func receive(cr, qr <-chan int, quit <-chan bool) {
	for {
		select {
		case cMsg, ok := <-cr:
			if ok {
				fmt.Println(cMsg)
			}
		case qMsg, ok := <-qr:
			if ok {
				fmt.Println(qMsg)
			}
		case _, ok := <-quit:
			if !ok {
				return
			}
		}
	}
}
