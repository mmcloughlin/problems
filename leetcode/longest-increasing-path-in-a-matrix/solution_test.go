package solution

import "testing"

func TestLongestIncreasingPathCases(t *testing.T) {
	cases := []struct {
		Matrix [][]int
		Expect int
	}{
		{
			Matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			Expect: 4,
		},
		{
			Matrix: [][]int{
				{9, 9, 9},
				{9, 9, 9},
				{9, 9, 9},
				{9, 9, 9},
			},
			Expect: 1,
		},
		{
			Matrix: [][]int{
				{9, 9, 9},
				{8, 8, 8},
				{7, 7, 7},
			},
			Expect: 3,
		},
	}
	for _, c := range cases {
		if got := longestIncreasingPath(c.Matrix); got != c.Expect {
			t.Fatalf("longestIncreasingPath(%v) = %v; expect %v", c.Matrix, got, c.Expect)
		}
	}
}
