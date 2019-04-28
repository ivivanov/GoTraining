package main

import "fmt"
import "time"

func main() {
	born := 1989
	today := time.Now().Year()
	i := born
	for i <= today {
		fmt.Println(i)
		i++
	}
}
