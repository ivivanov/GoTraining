package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func main() {
	html, _ := readFile("index.html")

	type person struct {
		ID   int
		Name string
		Age  int
	}

	data := []person{
		person{0, "Peter", 22},
		person{1, "Jo", 1},
		person{2, "Stef", 12},
		person{3, "Parker", 33},
		person{4, "Dr", 40},
		person{5, "Strange", 45},
	}

	tpl := template.Must(template.New("html").Parse(html))
	buf := bytes.NewBuffer([]byte{})
	tpl.Execute(buf, data)
	writeFile(buf.Bytes())
}

func readFile(path string) (string, error) {
	res, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func writeFile(data []byte) {
	ioutil.WriteFile("out.html", data, 0644)
}
