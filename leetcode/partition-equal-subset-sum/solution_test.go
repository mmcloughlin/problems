package solution

import "testing"

func TestCanPartition(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{5, 5}, true},
		{[]int{5, 6}, false},
		{[]int{6}, false},
		{[]int{}, false},
		{[]int{1, 2, 3, 5}, false},
	}
	for _, c := range cases {
		if got := canPartition(c.Input); got != c.Expect {
			t.Errorf("canPartition(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
