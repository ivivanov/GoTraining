package main

import "fmt"

var human map[string][]string

func main() {
	human = map[string][]string{
		"iv_ivanov": []string{"312", "asd"},
		"st_kirov":  []string{"000", "qwer"},
		"pe_petrov": []string{"111", "ger"}}

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
