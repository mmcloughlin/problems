package solution

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	cases := []struct {
		Input  []int
		Target int
		Expect [][]int
	}{
		{
			Input:  []int{2, 3, 6, 7},
			Target: 7,
			Expect: [][]int{
				{7},
				{2, 2, 3},
			},
		},
		{
			Input:  []int{2, 3, 5},
			Target: 8,
			Expect: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			Input:  []int{1},
			Target: 8,
			Expect: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 1},
			},
		},
		{
			Input:  []int{},
			Target: 8,
			Expect: nil,
		},
		{
			Input:  []int{},
			Target: 0,
			Expect: [][]int{{}},
		},
	}
	for _, c := range cases {
		if got := combinationSum(c.Input, c.Target); !combinationsEqual(got, c.Expect) {
			t.Errorf("combinationSum(%v, %d) = %v; expect %v", c.Input, c.Target, got, c.Expect)
		}
	}
}

func combinationsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for _, comb := range a {
		found := false
		for _, other := range b {
			if reflect.DeepEqual(comb, other) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
