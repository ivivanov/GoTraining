package main

import "fmt"

// declaration of human
var human map[string][]string

func main() {
	fmt.Printf("%T\n", human)

	// human initialization
	human = make(map[string][]string, 5)

	human["iv_ivanov"] = []string{"312", "asd"}
	human["pe_petrov"] = []string{"111", "ger"}
	human["st_kirov"] = []string{"000", "qwer"}

	for k, v := range human {
		fmt.Println("key:", k, "value:", v)
	}
}
