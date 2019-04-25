package solution

import "testing"

func TestLeastInterval(t *testing.T) {
	cases := []struct {
		Tasks  string
		N      int
		Expect int
	}{
		{"AAABBB", 2, 8},
		{"AAABBB", 3, 10},
		{"AAABBB", 100, 204},
		{"AAAAAA", 2, 16},
		{"ABCDEF", 2, 6},
		{"AAAAAA", 1, 11},
		{"AAAAAA", 0, 6},
		{"AAAAAABCDEFG", 2, 16},
	}
	for _, c := range cases {
		if got := leastInterval([]byte(c.Tasks), c.N); got != c.Expect {
			t.Errorf("leastInterval(%v, %d) = %v; expect %v", c.Tasks, c.N, got, c.Expect)
		}
	}
}
