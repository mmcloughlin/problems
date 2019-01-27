package search

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		Input  []int
		Target int
		Expect int
	}{
		{
			Input:  []int{4, 5, 6, 7, 0, 1, 2},
			Target: 0,
			Expect: 4,
		},
		{
			Input:  []int{4, 5, 6, 7, 0, 1, 2},
			Target: 3,
			Expect: -1,
		},
		{
			Input:  []int{4},
			Target: 4,
			Expect: 0,
		},
		{
			Input:  []int{4},
			Target: 5,
			Expect: -1,
		},
		{
			Input:  []int{0, 1, 2, 3, 4, 5},
			Target: 1,
			Expect: 1,
		},
		{
			Input:  []int{},
			Target: 5,
			Expect: -1,
		},
	}
	for _, c := range cases {
		got := search(c.Input, c.Target)
		if got != c.Expect {
			t.Errorf("search(%v, %v) = %v; expect %v", c.Input, c.Target, got, c.Expect)
		}
	}
}
