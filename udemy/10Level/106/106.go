package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rnd *rand.Rand

func main() {
	rndSeed := time.Now().UnixNano()
	rnd = rand.New(rand.NewSource(rndSeed))
	chanArr := make([]chan int, 10)

	// create multiple channels and fill them with data
	for i := 0; i < len(chanArr); i++ {
		chanArr[i] = make(chan int)

		go func(c chan int) {
			c <- rnd.Intn(100)
			close(c)
		}(chanArr[i])
	}
	//

	for v := range merge(chanArr...) {
		fmt.Println(v)
	}

	fmt.Println("Main before exit")
}

func merge(channels ...chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	out := make(chan int)

	for _, c := range channels {

		go func(c chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)

	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
