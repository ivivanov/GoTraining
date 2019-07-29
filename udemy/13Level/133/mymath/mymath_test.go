package mymath

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	type test struct {
		data []int
		ans  int
	}

	tests := []test{
		test{[]int{1, 1}, 2},
		test{[]int{1, -1}, 0},
		test{[]int{1, 1, 1}, 3},
		test{[]int{1, 0, 0}, 1},
	}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.ans {
			t.Error("Expected", v.ans, "got", x)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 3, -2))
	// Output:
	// 2
}
