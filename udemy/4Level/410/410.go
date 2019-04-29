package main

import "fmt"

// declaration of human
var human map[string][]string

func main() {
	fmt.Printf("%T\n", human)

	// human initialization
	human = make(map[string][]string, 3)

	human["iv_ivanov"] = []string{"312", "asd"}
	human["pe_petrov"] = []string{"111", "ger"}
	human["st_kirov"] = []string{"000", "qwer"}
	printMap()
	human["pe_petrov"] = human["pe_petrov"][1:] // leave only "ger" in the slice
	printMap()

	delete(human, "iv_ivanov")
	printMap()

}

func printMap() {
	for k, v := range human {
		fmt.Println("key:", k, "value:", v)
	}
	fmt.Println()
}
