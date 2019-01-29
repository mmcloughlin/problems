package ways

import "testing"

func TestNumDecodings(t *testing.T) {
	cases := []struct {
		Digits string
		Expect int
	}{
		{"12", 2},
		{"226", 3},
		{"1", 1},
		{"10", 1},
		{"0", 0},
		{"611", 2},
	}
	for _, c := range cases {
		got := numDecodings(c.Digits)
		if got != c.Expect {
			t.Errorf("numDecodings(%q) = %d; expect %d", c.Digits, got, c.Expect)
		}
	}
}
