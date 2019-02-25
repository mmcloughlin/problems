package solution

import "testing"

func TestStrStr(t *testing.T) {
	cases := []struct {
		Haystack string
		Needle   string
		Expect   int
	}{
		{"a", "hello", -1},
		{"hello", "hello", 0},
		{"hello", "", 0},
		{"hello", "e", 1},
		{"hello", "ll", 2},
		{"hello", "he", 0},
		{"hello", "el", 1},
		{"aaaaaaaa", "bba", -1},
	}
	for _, c := range cases {
		if got := strStr(c.Haystack, c.Needle); got != c.Expect {
			t.Errorf("strStr(%v, %v) = %v; expect %v", c.Haystack, c.Needle, got, c.Expect)
		}
	}
}
