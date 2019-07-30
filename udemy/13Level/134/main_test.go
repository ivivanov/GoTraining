package main

import "testing"

func Test(t *testing.T) {
	result := SayHi()
	if result != "Hi" {
		t.Error("Expected", "Hi", "got", result)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHi()
	}
}
