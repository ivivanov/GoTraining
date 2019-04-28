package main

import "fmt"
import "time"

func main() {
	born := 1989
	today := time.Now().Year()
	i := born
	for {
		if i > today {
			break
		}
		fmt.Println(i)
		i++
	}
}
