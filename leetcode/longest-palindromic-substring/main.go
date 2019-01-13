package main

import "log"

func longestPalindrome(s string) string {
	max := 0
	sub := ""
	n := len(s)
	for i := 0; i < n; i++ {
		for p := 0; p < 2; p++ {
			// invariant: s[l:l+size] is palendrome
			l, size := i, p
			for ; l > 0 && l+size < n && s[l-1] == s[l+size]; l, size = l-1, size+2 {
			}
			if size > max {
				max = size
				sub = s[l : l+size]
			}
		}
	}
	return sub
}

func main() {
	cases := []struct {
		Input  string
		Output string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"abcdef", "a"},
		{"", ""},
		{"aaaa", "aaaa"},
		{"aaaaa", "aaaaa"},
		{"baaaaa", "aaaaa"},
	}
	for _, c := range cases {
		got := longestPalindrome(c.Input)
		if got != c.Output {
			log.Fatal("input %q got %q expect %q", c.Input, got, c.Output)
		}
	}
	log.Print("pass")
}
