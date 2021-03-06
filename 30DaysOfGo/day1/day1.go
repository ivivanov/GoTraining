package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	num, _ := strconv.ParseUint(scanner.Text(), 10, 64)

	scanner.Scan()
	dub, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	str := scanner.Text()

	fmt.Println(i + num)
	fmt.Printf("%.1f", d+dub)
	fmt.Println()
	fmt.Println(s + str)
}
