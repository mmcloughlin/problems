package solution

import "bytes"

// A is the hash multiplier.
const A = 6364136223846793005

func strStr(haystack string, needle string) int {
	nb := []byte(needle)
	hb := []byte(haystack)

	n := len(nb)

	if len(hb) < n {
		return -1
	}
	if n == 0 {
		return 0
	}

	// Compute target hash.
	target := hash(nb)

	// Compute rolling hash through haystack.
	h := hash(hb[:n-1])
	An1 := pow(A, n-1)
	for i := 0; i+n <= len(hb); i++ {
		h = A*h + uint64(hb[i+n-1])

		if h == target && bytes.Equal(hb[i:i+n], nb) {
			return i
		}

		h -= An1 * uint64(hb[i])
	}

	return -1
}

func hash(s []byte) uint64 {
	h := uint64(0)
	for _, b := range s {
		h = A*h + uint64(b)
	}
	return h
}

func pow(x uint64, n int) uint64 {
	p := uint64(1)
	for ; n > 0; n-- {
		p *= x
	}
	return p
}
