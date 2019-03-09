package solution

func expressiveWords(S string, words []string) int {
	sr := rle(S)
	n := 0
	for _, w := range words {
		wr := rle(w)
		if stretchable(sr, wr) {
			n++
		}
	}
	return n
}

type run struct {
	ch byte
	n  int
}

// return true if w can be stretched into s.
func stretchable(s, w []run) bool {
	if len(s) != len(w) {
		return false
	}
	for i := range s {
		require := s[i].ch == w[i].ch && (s[i].n == w[i].n || (w[i].n < s[i].n && s[i].n >= 3))
		if !require {
			return false
		}
	}
	return true
}

func rle(w string) []run {
	runs := []run{}
	b := []byte(w)
	for i := 0; i < len(b); {
		s := i
		for ; i < len(b) && b[s] == b[i]; i++ {
		}
		runs = append(runs, run{ch: b[s], n: i - s})
	}
	return runs
}
