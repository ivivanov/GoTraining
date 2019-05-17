package main

import (
	"encoding/json"
	"fmt"
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

	bytes, err := json.Marshal(arr)

	if err != nil {
		fmt.Println("marshal:", err)
		return
	}

	fmt.Printf("%s", bytes)
}
