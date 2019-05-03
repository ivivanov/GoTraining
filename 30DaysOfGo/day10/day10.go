package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s, _, _ := r.ReadLine()
	n, _ := strconv.ParseInt(string(s), 10, 64)
	binary := fmt.Sprintf("%b", n)

	maxConsecutive := findMaxConsecutiveOnes(binary)
	fmt.Println(maxConsecutive)
}

func findMaxConsecutiveOnes(s string) int {
	arr := strings.Split(s, "0")
	return longestString(arr)
}

func longestString(a []string) int {
	maxLen := 0
	for _, v := range a {
		if maxLen < len(v) {
			maxLen = len(v)
		}
	}
	return maxLen
}
