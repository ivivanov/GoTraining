package main

import "fmt"

func main() {
	for i := 1; i < 10001; i++ {
		fmt.Printf("%v\t", i)
		if i%100 == 0 {
			fmt.Println()
		}
	}
}
