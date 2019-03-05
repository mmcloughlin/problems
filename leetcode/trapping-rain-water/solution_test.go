package solution

import "testing"

func TestTrap(t *testing.T) {
	cases := []struct {
		Height []int
		Expect int
	}{
		{
			Height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			Expect: 6,
		},
	}
	for _, c := range cases {
		if got := trap(c.Height); got != c.Expect {
			t.Errorf("trap(%v) = %v; expect %v", c.Height, got, c.Expect)
		}
	}
}
