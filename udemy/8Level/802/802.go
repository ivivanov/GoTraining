package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {

	text, err := ioutil.ReadFile("text.json")

	if err != nil {
		fmt.Println("read file:", err)
	}

	var personArr []person
	err = json.Unmarshal(text, &personArr)

	if err != nil {
		fmt.Println("unmarshal:", err)
	}

	fmt.Println(personArr)
}
