package main

import "log"

func lengthOfLongestSubstring(s string) int {
	max := 0
	seen := map[byte]bool{}
	t := 0
	for i := range s {
		c := s[i]

		// If this is a new character, add it to the seen set and (possibly)
		// update the max length.
		if !seen[c] {
			seen[c] = true
			if len(seen) > max {
				max = len(seen)
			}
			continue
		}

		// Otherwise we have seen the character before, so we need to advance the
		// tail pointer until we find the repeat.
		for ; s[t] != c; t++ {
			delete(seen, s[t])
		}
		t++
	}
	return max
}

func main() {
	cases := []struct {
		Input  string
		Expect int
	}{
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},
		{"aaaaaaaaaaaaabbbbbb", 2},
		{"asdfghjklqwertyuiop", 19},
		{"", 0},
	}
	for _, c := range cases {
		got := lengthOfLongestSubstring(c.Input)
		if got != c.Expect {
			log.Fatalf("error on %q", c.Input)
		}
		log.Printf("%q result %d", c.Input, got)
	}
	log.Print("pass")
}
