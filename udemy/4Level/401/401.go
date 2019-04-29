package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 11
	arr[1] = 12
	arr[2] = 13
	arr[3] = 14
	arr[4] = 15

	for i, v := range arr {
		fmt.Printf("index %v : value %v\n", i, v)
	}

	fmt.Printf("%T\n", arr)
	fmt.Println(len(arr), cap(arr), arr)
}
