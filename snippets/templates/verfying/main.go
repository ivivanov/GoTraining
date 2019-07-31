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

// Saying is template for describing a person
const Saying string = "We call him '{{ .Name }}' and he is {{ .Age }} years old."

func main() {
	person := Person{"Ivan", 29}
	tpl, _ := ParseSayingTpl()
	byteString, _ := ExecuteSayingTpl(&person, tpl)
	fmt.Println(string(byteString))
}

// ParseSayingTpl creates template based on Saying
func ParseSayingTpl() (*template.Template, error) {
	tpl, err := template.New("saying").Parse(Saying)
	if err != nil {
		fmt.Println("Error during text template parsing", err)
		return nil, err
	}

	return tpl, nil
}

// ExecuteSayingTpl matches person with templated based on person
func ExecuteSayingTpl(p *Person, tpl *template.Template) ([]byte, error) {
	writer := bytes.NewBuffer([]byte{})
	err := tpl.Execute(writer, p)
	if err != nil {
		fmt.Println("Error during template execution", err)
		return nil, err
	}

	return writer.Bytes(), nil
}
