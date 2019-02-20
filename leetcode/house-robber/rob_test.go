package rob

import "testing"

func TestRobHouses(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{100, 1, 1, 100}, 200},
	}
	for _, c := range cases {
		if got := rob(c.Input); got != c.Expect {
			t.Errorf("robHouses(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
