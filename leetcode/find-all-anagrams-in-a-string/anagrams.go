package anagrams

func findAnagrams(s, p string) []int {
	// Character counts in p.
	counts := map[byte]int{}
	for _, ch := range []byte(p) {
		counts[ch]++
	}

	b := []byte(s)
	n := len(p)
	idxs := []int{}
	for i, ch := range b {
		// Update counts.
		add(counts, ch, -1)
		s := i - n
		if s >= 0 {
			add(counts, b[s], +1)
		}

		if len(counts) == 0 {
			idxs = append(idxs, s+1)
		}
	}

	return idxs
}

func add(counts map[byte]int, ch byte, i int) {
	counts[ch] += i
	if counts[ch] == 0 {
		delete(counts, ch)
	}
}
