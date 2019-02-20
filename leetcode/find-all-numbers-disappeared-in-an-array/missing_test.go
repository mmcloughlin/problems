package missing

import (
	"reflect"
	"testing"
)

func TestFindDisappeared(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect []int
	}{
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{},
		},
		{
			[]int{1, 2, 2, 4, 5, 6},
			[]int{3},
		},
	}
	for _, c := range cases {
		if got := findDisappearedNumbers(c.Input); !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("findDisappeared(%v) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
