package solution

import "math"

func firstUniqChar(s string) int {
	first := map[byte]int{}
	count := map[byte]int{}
	b := []byte(s)
	for i := len(b) - 1; i >= 0; i-- {
		first[b[i]] = i
		count[b[i]]++
	}

	idx := math.MaxInt64
	found := false
	for ch, n := range count {
		if n == 1 && first[ch] < idx {
			found = true
			idx = first[ch]
		}
	}

	if !found {
		return -1
	}
	return idx
}
