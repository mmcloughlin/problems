package ways

func numDecodings(s string) int {
	num := 1
	pre := 0
	n := 0
	for _, b := range []byte(s) {
		d := int(b - '0')
		n = (10*n + d) % 100

		nxt := 0
		if d > 0 {
			nxt += num
		}
		if 10 <= n && n <= 26 {
			nxt += pre
		}

		pre, num = num, nxt
	}

	return num
}
