package math

import "testing"

func TestSum(t *testing.T) {
	x := Sum(1, 2)

	if x != 3 {
		t.Error("Expected:", 3, " Got:", x)
	}
}
