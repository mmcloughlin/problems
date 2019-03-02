package solution

import "strings"

func longestSubstring(s string, k int) int {
	b := []byte(s)
	if len(b) == 0 {
		return 0
	}

	// Compute character counts.
	count := map[byte]int{}
	for _, ch := range b {
		count[ch]++
	}

	// If they are all at least k, we're done.
	allatleastk := true
	var split byte
	for ch, c := range count {
		if c < k {
			allatleastk = false
			split = ch
			break
		}
	}
	if allatleastk {
		return len(b)
	}

	// Otherwise, split on the character and recurse.
	parts := strings.Split(s, string([]byte{split}))
	m := 0
	for _, part := range parts {
		n := longestSubstring(part, k)
		if n > m {
			m = n
		}
	}

	return m
}
