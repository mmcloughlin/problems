package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect [][]int
	}{
		{
			Input: []int{-1, 0, 1, 2, -1, -4},
			Expect: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			Input: []int{-1, -1, -1, 0, 1, 1, 1},
			Expect: [][]int{
				{-1, 0, 1},
			},
		},
		{
			Input:  []int{1, 2, -2, -1},
			Expect: [][]int{},
		},
	}
	for _, c := range cases {
		got := threeSum(c.Input)
		if !reflect.DeepEqual(got, c.Expect) {
			t.Logf("got\n%#v", got)
			t.Logf("expect\n%#v", c.Expect)
			t.FailNow()
		}
	}
}
