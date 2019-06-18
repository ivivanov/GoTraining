package main

import (
	"fmt"

	"github.com/ivivanov/GoTraining/udemy/11Level/tests/hmm"
)

func main() {
	var i int
	foo(&i)
	fmt.Println(i)
	hmm.DoSth()
}

func foo(i *int) {
	defer func(x *int) {
		fmt.Println("Defer func")
		*x++
	}(i)

	fmt.Println("before ret")
}
