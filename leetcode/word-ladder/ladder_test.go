package ladder

import "testing"

func TestLadderLength(t *testing.T) {
	cases := []struct {
		Begin  string
		End    string
		Words  []string
		Expect int
	}{
		{
			Begin:  "hit",
			End:    "cog",
			Words:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			Expect: 5,
		},
	}
	for _, c := range cases {
		got := ladderLength(c.Begin, c.End, c.Words)
		if got != c.Expect {
			t.Errorf("ladderLength(%q, %q, %s) = %d; expect %d", c.Begin, c.End, c.Words, got, c.Expect)
		}
	}
}
