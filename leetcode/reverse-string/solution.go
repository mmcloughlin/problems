package solution

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[r], s[l] = s[l], s[r]
	}
}
