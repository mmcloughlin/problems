package solution

import "testing"

func TestCountAndSay(t *testing.T) {
	expect := []string{
		"1",
		"11",
		"21",
		"1211",
		"111221",
		"312211",
	}
	for n, e := range expect {
		if got := countAndSay(n + 1); got != e {
			t.Errorf("countAndSay(%v) = %v; expect %v", n+1, got, e)
		}
	}
}
