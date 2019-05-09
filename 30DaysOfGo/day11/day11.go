package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
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
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	var biggestClock = int32(math.MinInt32)

	for i := 0; i < size-maxI; i++ {
		for j := 0; j < size-maxJ; j++ {
			currentSum := clockSum(i, j, arr)
			if biggestClock < currentSum {
				biggestClock = currentSum
			}
		}
	}

	fmt.Println(biggestClock)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// ai, aj - top left corner coordinates
func clockSum(ai, aj int, matrix [][]int32) int32 {
	// validate input. Clock should not go out matrix boundaries
	if ai < 0 || aj < 0 || ai+maxI >= size || aj+maxJ >= size {
		panic("incorrect starting coordinates")
	}

	var sum int32
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
