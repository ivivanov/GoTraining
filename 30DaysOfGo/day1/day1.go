package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	// var i uint64 = 4
	// var d float64 = 4.0
	// var s string = "HackerRank"

	var (
		num int64
		dub float64
		str string
	)

	r := bufio.NewReader(os.Stdin)
	for i := 0; i < 3; i++ {
		fmt.Fscanln(r, "%d %b %s", &num, &dub, &str)
	}

	fmt.Println(num)
	fmt.Println(dub)
	fmt.Println(str)
	// scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

}