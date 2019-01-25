package zero

import (
	"reflect"
	"testing"
)

func TestSetZeros(t *testing.T) {
	cases := []struct {
		Matrix [][]int
		Expect [][]int
	}{
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, c := range cases {
		setZeroes(c.Matrix)
		if !reflect.DeepEqual(c.Matrix, c.Expect) {
			t.Log(c.Matrix)
			t.Log(c.Expect)
			t.Fail()
		}
	}
}
