package schedule

import "testing"

func TestCanFinish(t *testing.T) {
	cases := []struct {
		Num           int
		Prerequisites [][]int
		Expect        bool
	}{
		{
			Num: 2,
			Prerequisites: [][]int{
				{1, 0},
			},
			Expect: true,
		},
		{
			Num: 2,
			Prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			Expect: false,
		},
		{
			Num: 4,
			Prerequisites: [][]int{
				{1, 0},
				{2, 1},
				{3, 2},
				{0, 3},
			},
			Expect: false,
		},
		{
			Num:           400,
			Prerequisites: [][]int{},
			Expect:        true,
		},
	}
	for _, c := range cases {
		got := canFinish(c.Num, c.Prerequisites)
		if got != c.Expect {
			t.Errorf("canFinish(%v, %v) = %v; expect %v", c.Num, c.Prerequisites, got, c.Expect)
		}
	}
}
