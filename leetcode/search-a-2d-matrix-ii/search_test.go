package box

import "testing"

var example = [][]int{
	{1, 4, 7, 11, 15},
	{2, 5, 8, 12, 19},
	{3, 6, 9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30},
}

func TestSearchMatrix(t *testing.T) {
	example := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	cases := []struct {
		Matrix [][]int
		Target int
		Expect bool
	}{
		{example, 5, true},
		{example, 20, false},
	}
	for _, c := range cases {
		got := searchMatrix(c.Matrix, c.Target)
		if got != c.Expect {
			t.Errorf("searchMatrix(%v, %v) = %v; expect %v", c.Matrix, c.Target, got, c.Expect)
		}
	}
}

func TestSearchAll(t *testing.T) {
	for _, row := range example {
		for _, target := range row {
			if !searchMatrix(example, target) {
				t.Errorf("could not find %v", target)
			}
		}
	}
}
