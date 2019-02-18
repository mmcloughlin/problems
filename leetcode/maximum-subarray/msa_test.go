package msa

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := 6
	if got := maxSubArray(nums); got != expect {
		t.Errorf("maxSubArray(%v) = %v; expect %v", nums, got, expect)
	}
}
