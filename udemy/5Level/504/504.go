package main

import "fmt"

func main() {
	anon := struct {
		pages      int
		size       string
		name       string
		characters map[string][]string
	}{
		pages: 2523,
		size:  "A4",
		name:  "Lord of the rings",
		characters: map[string][]string{
			"orcs":  []string{"Azog", "Grishnákh", "Muzgash"},
			"elves": []string{"Tauriel", "Gil-Galad", "FËANOR"}}}

	fmt.Println(anon)
}
