package solution

func titleToNumber(s string) int {
	n := 0
	for _, b := range []byte(s) {
		v := b - 'A' + 1
		n = 26*n + int(v)
	}
	return n
}
