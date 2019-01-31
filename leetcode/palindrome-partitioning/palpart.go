package palpart

func partition(s string) [][]string {
	n := len(s)

	// possibleAt[i] is the possible partitions of s[:i]
	possibleAt := make([][][]string, n+1)
	possibleAt[0] = [][]string{
		{},
	}

	for i := 0; i < n; i++ {
		spans := palendromesCenteredAt(s, i)

		for _, spn := range spans {
			for _, pos := range possibleAt[spn.Left] {
				newsplit := []string{}
				newsplit = append(newsplit, pos...)
				newsplit = append(newsplit, s[spn.Left:spn.Right])
				possibleAt[spn.Right] = append(possibleAt[spn.Right], newsplit)
			}
		}
	}

	return possibleAt[n]
}

type span struct {
	Left  int
	Right int
}

func palendromesCenteredAt(s string, i int) []span {
	spans := []span{}

	// Odd length.
	l, r := i, i
	for ; 0 <= l && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
		spans = append(spans, span{l, r + 1})
	}

	// Even length.
	l, r = i, i+1
	for ; 0 <= l && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
		spans = append(spans, span{l, r + 1})
	}

	return spans
}
