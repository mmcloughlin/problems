package solution

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		Nums   []int
		K      int
		Expect []int
	}{
		{
			Nums:   []int{1, 1, 1, 2, 2, 3},
			K:      2,
			Expect: []int{1, 2},
		},
		{
			Nums:   []int{1},
			K:      1,
			Expect: []int{1},
		},
	}
	for _, c := range cases {
		if got := topKFrequent(c.Nums, c.K); !reflect.DeepEqual(got, c.Expect) {
			t.Fatalf("topKFrequent(%v, %v) = %v; expect %v", c.Nums, c.K, got, c.Expect)
		}
	}
}
