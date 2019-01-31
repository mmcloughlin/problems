package gas

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	cases := []struct {
		Gas    []int
		Cost   []int
		Expect int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
		{
			[]int{3, 3, 4},
			[]int{3, 4, 4},
			-1,
		},
	}
	for _, c := range cases {
		got := canCompleteCircuit(c.Gas, c.Cost)
		if got != c.Expect {
			t.Errorf("canCompleteCircuit(%v, %v) = %v; expect %v", c.Gas, c.Cost, got, c.Expect)
		}
	}
}
