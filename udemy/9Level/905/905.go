package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

const c = 300

func main() {
	runtime.GOMAXPROCS(10)
	wg.Add(c)

	for i := 0; i < c; i++ {
		go increment(&counter)
	}

	wg.Wait()
	fmt.Println(counter)
}

func increment(counter *int64) {
	atomic.AddInt64(counter, 1)
	fmt.Println("addOne called ", atomic.LoadInt64(counter))
	wg.Done()
}
