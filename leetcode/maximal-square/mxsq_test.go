package mxsq

import "testing"

func TestMaximalSquare(t *testing.T) {
	cases := []struct {
		Matrix [][]byte
		Expect int
	}{
		{
			Matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			Expect: 4,
		},
		{
			Matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '0', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			Expect: 1,
		},
		{
			Matrix: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			Expect: 0,
		},
	}
	for _, c := range cases {
		got := maximalSquare(c.Matrix)
		if got != c.Expect {
			t.Errorf("maximalSquare(%v) = %v; expect %v", c.Matrix, got, c.Expect)
		}
	}
}
