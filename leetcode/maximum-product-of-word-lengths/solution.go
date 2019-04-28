package solution

func maxProduct(words []string) int {
	type word struct {
		n    int
		mask uint32
	}
	masks := []word{}
	for _, w := range words {
		masks = append(masks, word{
			n:    len(w),
			mask: mask(w),
		})
	}

	p := 0
	for _, w0 := range masks {
		for _, w1 := range masks {
			if (w0.mask&w1.mask) == 0 && w0.n*w1.n > p {
				p = w0.n * w1.n
			}
		}
	}

	return p
}

func mask(word string) uint32 {
	var m uint32
	for _, ch := range []byte(word) {
		m |= 1 << uint(ch-'a')
	}
	return m
}
