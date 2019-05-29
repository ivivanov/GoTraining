package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// runtime.GOMAXPROCS(3)
	wg.Add(2)

	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println(i, "foo")
			runtime.Gosched()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println(i, "bar")
			runtime.Gosched()
		}
		wg.Done()
	}()

	for i := 0; i < 30; i++ {
		fmt.Println(i, "main")
		runtime.Gosched()
	}

	wg.Wait()
}
