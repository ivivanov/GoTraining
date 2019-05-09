package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const size int = 6

//sandclock walls coordinates
//00,01,02
//   11
//20,21,22

type coordinates struct {
	i, j int
}

var clockWalls = [7]coordinates{
	coordinates{i: 0, j: 0},
	coordinates{i: 0, j: 1},
	coordinates{i: 0, j: 2},
	coordinates{i: 1, j: 1},
	coordinates{i: 2, j: 0},
	coordinates{i: 2, j: 1},
	coordinates{i: 2, j: 2},
}

var maxI = clockWalls[6].i
var maxJ = clockWalls[6].j

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	matrix := buildMatrix(string(input))
	// clockA := clockSum(3, 4, matrix)

	var biggestClock = math.MinInt32

	for i := 0; i < size-maxI; i++ {
		for j := 0; j < size-maxJ; j++ {
			currentSum := clockSum(i, j, matrix)
			if biggestClock < currentSum {
				biggestClock = currentSum
			}
		}
	}

	fmt.Println(biggestClock)
}

// ai, aj - top left corner coordinates
func clockSum(ai, aj int, matrix [size][size]int) int {
	// validate input. Clock should not go out matrix boundaries
	if ai < 0 || aj < 0 || ai+maxI >= size || aj+maxJ >= size {
		panic("incorrect starting coordinates")
	}

	var sum int
	for _, v := range clockWalls {
		sum += matrix[ai+v.i][aj+v.j]
	}

	return sum
}

func buildMatrix(s string) [size][size]int {
	reader := strings.NewReader(s)
	r := bufio.NewReader(reader)
	var matrix [size][size]int

	for i := 0; i < size; i++ {

		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		// fmt.Println(string(line))
		matrix[i] = convertLineToArray(string(line))
	}

	return matrix
}

func convertLineToArray(line string) [size]int {
	arr := strings.Split(line, " ")
	var result [size]int

	for i := 0; i < len(arr); i++ {
		lineItem, _ := strconv.ParseInt(arr[i], 10, 64)
		result[i] = int(lineItem)
	}

	return result
}
