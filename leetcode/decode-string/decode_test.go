package decode

import "testing"

func TestDecodeString(t *testing.T) {
	cases := []struct {
		Input  string
		Expect string
	}{
		{"abc", "abc"},
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, c := range cases {
		got := decodeString(c.Input)
		if got != c.Expect {
			t.Errorf("decodeString(%q) = %v; expect %v", c.Input, got, c.Expect)
		}
	}
}
