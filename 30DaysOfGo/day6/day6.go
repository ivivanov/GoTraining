package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	tStr := r.Text()
	t, _ := strconv.ParseInt(tStr, 10, 64)

	for i := 0; i < int(t); i++ {
		r.Scan()
		tmp := r.Text()
		even := make([]byte, 0)
		odd := make([]byte, 0)
		for j := 0; j < len(tmp); j++ {
			if j == 0 || j%2 == 0 {
				even = append(even, tmp[j])
			} else {
				odd = append(odd, tmp[j])
			}
		}
		fmt.Println(string(even), string(odd))
	}
}
