package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// Person is a struct representing a person
type Person struct {
	Name string
	Age  int
}

const saying string = "We call him '{{ .Name }}' and he is {{ .Age }} years old."

func main() {
	person := Person{"Ivan", 29}
	tpl, err := template.New("ivan").Parse(saying)
	if err != nil {
		fmt.Println("Error during text template parsing", err)
		panic(err)
	}

	writer := bytes.NewBuffer([]byte{})
	err = tpl.Execute(writer, person)
	if err != nil {
		fmt.Println("Error during template execution", err)
		panic(err)
	}

	out := string(writer.Bytes())
	fmt.Println(out)
}
