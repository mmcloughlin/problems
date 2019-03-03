package solution

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 1, 1, 1, 1, 2, 4}, true},
		{[]int{2, 7, 6, 5, 5, 5, 3}, false},
	}
	for _, c := range cases {
		if got := increasingTriplet(c.Input); got != c.Expect {
			t.Errorf("increasingTriplet(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
