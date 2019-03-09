package solution

import "testing"

func TestMajorityElement(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, c := range cases {
		if got := majorityElement(c.Input); got != c.Expect {
			t.Errorf("majorityElement(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
