package solution

import "testing"

func TestMinPathSum(t *testing.T) {
	cases := []struct {
		Input  [][]int
		Expect int
	}{
		{
			Input: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 1, 1},
			},
			Expect: 7,
		},
		{
			Input: [][]int{
				{1, 3, 100},
			},
			Expect: 104,
		},
		{
			Input: [][]int{
				{1},
				{3},
				{100},
			},
			Expect: 104,
		},
	}
	for _, c := range cases {
		t.Log(c.Input)
		if got := minPathSum(c.Input); got != c.Expect {
			t.Log(c.Input)
			t.Errorf("minPathSum(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
