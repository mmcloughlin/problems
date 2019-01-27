package jump

import "testing"

func TestCanJump(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{0, 0}, false},
		{[]int{1, 1, 1, 1, 1, 0}, true},
	}
	for _, c := range cases {
		got := canJump(c.Input)
		if got != c.Expect {
			t.Errorf("canJump(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
