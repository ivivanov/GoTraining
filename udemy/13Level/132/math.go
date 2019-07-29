package math

// Sum returns the sum of all integers in the slice ints
func Sum(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
