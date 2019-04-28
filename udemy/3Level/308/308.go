package main

import "fmt"

func main() {
	switch { //by default the expression is true
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("case true")
	case 1 == 1:
		fmt.Println("1==1")
	default:
		fmt.Println("case default")
	}
}
