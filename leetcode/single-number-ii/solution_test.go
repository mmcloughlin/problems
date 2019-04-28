package solution

import "testing"

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		Nums   []int
		Expect int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}
	for _, c := range cases {
		if got := singleNumber(c.Nums); got != c.Expect {
			t.Errorf("singleNumber(%v) = %v; expect %v", c.Nums, got, c.Expect)
		}
	}
}
