package solution

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}
	for _, c := range cases {
		if got := missingNumber(c.Input); got != c.Expect {
			t.Errorf("missingNumber(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
