package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name   string
	Scores []int
}

// override String
func (p person) String() string {
	return "1"
}

func main() {
	arr := []person{
		person{Name: "Ken", Scores: []int{4, 5, 6}},
		person{Name: "Rob", Scores: []int{4, 3, 4}},
		person{Name: "Pik", Scores: []int{2, 6, 6}},
	}

	err := json.NewEncoder(os.Stdout).Encode(arr)

	if err != nil {
		fmt.Println("encode:", err)
	}
}
