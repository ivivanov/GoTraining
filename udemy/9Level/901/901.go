package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i, "foo")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i, "bar")
		}
		wg.Done()
	}()

	fmt.Println("main")
	wg.Wait()
}
