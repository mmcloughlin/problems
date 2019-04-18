package solution

import "testing"

func TestLargestRectangleAreaCases(t *testing.T) {
	cases := []struct {
		Heights []int
		Expect  int
	}{
		{Heights: []int{2, 1, 5, 6, 2, 3}, Expect: 10},
		{Heights: []int{1, 1, 1, 1, 1, 6, 1, 1, 1, 1}, Expect: 10},
		{Heights: []int{1, 1, 1, 1, 1, 60, 1, 1, 1, 1}, Expect: 60},
		{Heights: []int{1}, Expect: 1},
		{Heights: []int{11}, Expect: 11},
		{Heights: []int{}, Expect: 0},
		{Heights: []int{2, 1, 2}, Expect: 3},
	}
	for _, c := range cases {
		if got := largestRectangleArea(c.Heights); got != c.Expect {
			t.Errorf("largestRectangleArea(%v) = %v; expect %v", c.Heights, got, c.Expect)
		}
	}
}
