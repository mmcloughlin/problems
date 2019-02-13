package calc

import "testing"

func TestCalculate(t *testing.T) {
	cases := []struct {
		Expr  string
		Value int
	}{
		{"2", 2},
		{"    2", 2},
		{"2   ", 2},
		{" 2  ", 2},
		{" 211111  ", 211111},
		{"3+2*2", 7},
		{" 3/2 ", 1},
		{" 3+5 / 2", 5},
		{"0", 0},
		{"3-3", 0},
		{"3+3", 6},
		{"3*3", 9},
		{"3/3", 1},
		{"1-1+1", 1},
	}
	for _, c := range cases {
		got := calculate(c.Expr)
		t.Logf("%q = %d", c.Expr, got)
		if got != c.Value {
			t.Errorf("expected %d", c.Value)
		}
	}
}
