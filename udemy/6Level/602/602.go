package main

import "fmt"

func main() {
	//v1 declare everything inline
	fmt.Println(foo([]int{1, 2, 3, 4}...))

	//v2
	arr := []int{1, 2, 3, 4}
	sum := foo(arr...)
	fmt.Println(sum)

	//v3
	var arr2 [4]int
	for i := range arr2 {
		arr2[i] = i + 1
	}

	fmt.Println(foo(arr2[:]...)) // convert array to slice

	// arr and slice
	var array [2]int
	array[0] = 4
	array[1] = 5

	fmt.Printf("%T : %v\n", array, array)
	fmt.Printf("%T : %v\n", array[:], array[:])
}

// variadic functions accepts only slices (and not arrays)
func foo(a ...int) int {
	var sum int

	for _, v := range a {
		sum += v
	}

	return sum
}
