package solution

func countSubstrings(s string) int {
	total := 0
	b := []byte(s)
	for i := range b {
		total += countSubstringsAt(b, i)
	}
	return total
}

func countSubstringsAt(b []byte, i int) int {
	n := len(b)
	count := 0

	for p := 0; p < 2; p++ {
		for l, r := i-p, i; 0 <= l && r < n && b[l] == b[r]; l, r = l-1, r+1 {
			count++
		}
	}

	return count
}
