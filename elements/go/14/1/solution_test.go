package solution

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		A, B, Expect []int
	}{
		{
			A:      []int{2, 3, 3, 5, 5, 6, 7, 7, 8, 12},
			B:      []int{5, 5, 6, 8, 8, 9, 10, 10},
			Expect: []int{5, 6, 8},
		},
		{
			A:      []int{2, 2, 2, 2, 2},
			B:      []int{2},
			Expect: []int{2},
		},
		{
			A:      []int{1, 2, 3, 4, 5, 6},
			B:      []int{2, 4, 6, 6, 6, 6, 6},
			Expect: []int{2, 4, 6},
		},
		{
			A:      []int{2, 3},
			B:      []int{},
			Expect: []int{},
		},
		{
			A:      []int{},
			B:      []int{2, 3},
			Expect: []int{},
		},
		{
			A:      []int{},
			B:      []int{},
			Expect: []int{},
		},
	}
	for _, c := range cases {
		if got := intersect(c.A, c.B); !reflect.DeepEqual(got, c.Expect) {
			t.Errorf("intersect(%v, %v) = %v; expect %v", c.A, c.B, got, c.Expect)
		}
	}
}
