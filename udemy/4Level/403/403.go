package main

import "fmt"

func main() {
	sl := []int{40, 41, 42, 43, 44, 45, 46}
	fmt.Println(sl)
	fmt.Println(sl[1:])
	fmt.Println(sl[:4])
	fmt.Println(sl[2:4])

}
