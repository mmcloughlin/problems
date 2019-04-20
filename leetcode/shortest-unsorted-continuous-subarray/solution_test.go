package solution

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect int
	}{
		{Input: []int{2, 6, 4, 8, 10, 9, 15}, Expect: 5},
		{Input: []int{1, 2, 3, 4, 5}, Expect: 0},
		{Input: []int{4, 3, 2, 1}, Expect: 4},
		{Input: []int{1}, Expect: 0},
	}
	for _, c := range cases {
		if got := findUnsortedSubarray(c.Input); got != c.Expect {
			t.Fatalf("findUnsortedSubarray(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
