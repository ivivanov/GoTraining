package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup
var rnd *rand.Rand

const c = 300

func main() {
	rnd = rand.New(rand.NewSource(3))
	wg.Add(c)
	runtime.GOMAXPROCS(3)

	for i := 0; i < c; i++ {
		go increment(&counter)
	}

	wg.Wait()
	fmt.Println(counter)
}

func increment(i *int) {
	sleepTime := time.Second * time.Duration(rnd.Intn(5))
	// start some work
	time.Sleep(sleepTime)
	tmp := *i
	tmp++
	*i = tmp
	// end some work
	fmt.Println("addOne called ", *i)
	runtime.Gosched()
	wg.Done()
}
