package main

import "fmt"

func main() {
	slice := []int{312, 4, 1, 41, 41, 546, 12, 36, 76, 12}
	fmt.Printf("%T", slice)
	for i, v := range slice {
		fmt.Printf("%02v\t%v\n", i, v)
	}
}
