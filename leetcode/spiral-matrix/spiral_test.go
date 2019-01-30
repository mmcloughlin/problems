package spiral

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	cases := []struct {
		Matrix [][]int
		Expect []int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				{1, 2, 3},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 9, 8, 7},
		},
		{
			[][]int{{1}},
			[]int{1},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, c := range cases {
		got := spiralOrder(c.Matrix)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("\n   got=%v\nexpect=%v", got, c.Expect)
		}
	}
}
