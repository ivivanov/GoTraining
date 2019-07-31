package main

import (
	"testing"
)

func TestParseSayingTpl(t *testing.T) {
	_, err := ParseSayingTpl()
	if err != nil {
		t.Error(err)
	}
}

// Verifies that Saying template is matching Person struct
func TestExecuteSayingTpl(t *testing.T) {
	person := Person{"Ivan", 29}
	tpl, _ := ParseSayingTpl()
	_, err := ExecuteSayingTpl(&person, tpl)
	if err != nil {
		t.Error(err)
	}
}

// Verifies that ExecuteSayingTpl is producing valid result
func TestExecuteSayingTplForValidResult(t *testing.T) {
	person := Person{"Ivan", 29}
	expected := "We call him 'Ivan' and he is 29 years old."
	tpl, _ := ParseSayingTpl()
	byteString, _ := ExecuteSayingTpl(&person, tpl)
	if string(byteString) != expected {
		t.Error("Expected", expected, "got", string(byteString))
	}
}
