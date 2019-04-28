package solution

import "testing"

func TestNthUglyNumber(t *testing.T) {
	sequence := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12}
	for i, expect := range sequence {
		n := i + 1
		if got := nthUglyNumber(n); got != expect {
			t.Errorf("nthUglyNumber(%d) = %d; expect %d", n, got, expect)
		}
	}
}
