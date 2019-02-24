package jewels

func numJewelsInStones(J string, S string) int {
	// Build jewels mask.
	var jewels uint64
	for _, j := range []byte(J) {
		jewels |= uint64(1) << lettercode(j)
	}

	// Count jewels.
	n := 0
	for _, s := range []byte(S) {
		if (jewels & (uint64(1) << lettercode(s))) != 0 {
			n++
		}
	}

	return n
}

func lettercode(b byte) uint {
	if b < 'a' {
		return uint(b - 'A')
	}
	return uint(b-'a') + 32
}
