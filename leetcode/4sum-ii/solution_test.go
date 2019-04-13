package solution

import "testing"

func TestFourSumCountExample(t *testing.T) {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	n := fourSumCount(A, B, C, D)
	if n != 2 {
		t.Fail()
	}
}
