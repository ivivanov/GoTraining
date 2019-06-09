package main

import "fmt"

func main() {
	count := 3
	songChan := singSong("Tod")

	for i := 0; i < count; i++ {
		v, ok := <-songChan
		if ok {
			fmt.Println(v, ok)
		}
	}
}

func singSong(name string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprint("Happy birthday ", name)
		}
	}()

	return c
}
