package largest

import "testing"

func TestLargestNumber(t *testing.T) {
	cases := []struct {
		Nums   []int
		Expect string
	}{
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{34, 3456}, "345634"},
		{[]int{34, 3412}, "343412"},
		{[]int{3456, 34}, "345634"},
		{[]int{3412, 34}, "343412"},
		{[]int{3, 3}, "33"},
		{[]int{0, 0}, "0"},
		{[]int{0}, "0"},
		{[]int{0, 1, 0}, "100"},
	}
	for _, c := range cases {
		got := largestNumber(c.Nums)
		if got != c.Expect {
			t.Errorf("largestNumber(%v) = %v; expect %v", c.Nums, got, c.Expect)
		}
	}
}
