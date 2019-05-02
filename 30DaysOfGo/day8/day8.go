package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _ := strconv.ParseInt(read(reader), 10, 64)

	phoneBook := make(map[string]string, n)

	for i := 0; i < int(n); i++ {
		namePhoneEntry := read(reader)
		splitted := strings.Split(namePhoneEntry, " ")
		phoneBook[splitted[0]] = splitted[1]
	}

	for {
		query := read(reader)
		if query == "" {
			break
		}

		if res := phoneBook[query]; res == "" {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%v=%v\n", query, res)
		}
	}
}

func read(reader *bufio.Reader) string {
	line, _, err := reader.ReadLine()

	if err == io.EOF {
		return ""
	}

	return string(line)
}
