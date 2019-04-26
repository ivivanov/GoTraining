package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World.")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
