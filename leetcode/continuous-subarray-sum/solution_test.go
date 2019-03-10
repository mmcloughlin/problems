package solution

import (
	"testing"
)

func TestCheckSubarraySum(t *testing.T) {
	cases := []struct {
		Input  []int
		K      int
		Expect bool
	}{
		{
			Input:  []int{23, 2, 4, 6, 7},
			K:      6,
			Expect: true,
		},
		{
			Input:  []int{23, 2, 6, 4, 7},
			K:      6,
			Expect: true,
		},
		{
			Input:  []int{23, 2, 6, 4, 7},
			K:      0,
			Expect: false,
		},
		{
			Input:  []int{23, 2, 0, 0, 7},
			K:      0,
			Expect: true,
		},
		{
			Input:  []int{23, 2, 2, 0, 7},
			K:      0,
			Expect: false,
		},
	}
	for _, c := range cases {
		got := checkSubarraySum(c.Input, c.K)
		if got != c.Expect {
			t.Errorf("checkSubarraySum(%v, %v) = %v; expext %v", c.Input, c.K, got, c.Expect)
		}
	}
}
