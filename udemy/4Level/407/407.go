package main

import "fmt"

func main() {
	matrix := [][]string{
		[]string{"00", "01", "02"},
		[]string{"10", "11", "12"},
		[]string{"20", "21", "22"}}

	printMatrix(matrix)
	printMatrix2(matrix)
}

func printMatrix(matrix [][]string) {
	for i := 0; i < len(matrix[0]); i++ {
		fmt.Println()
		for j := 0; j < len(matrix[1]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
	}
}

func printMatrix2(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println()
		for _, v := range row {
			fmt.Printf("%v ", v)
		}
	}
}
