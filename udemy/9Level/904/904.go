package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

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

func increment(i *int) {
	mutex.Lock()
	tmp := *i
	tmp++
	*i = tmp
	fmt.Println("addOne called ", *i)
	mutex.Unlock()
	wg.Done()
}
