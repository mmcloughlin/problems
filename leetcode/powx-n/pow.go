package pow

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	p := 1.0
	for n != 0 {
		if (n & 1) == 1 {
			p *= x
		}
		x *= x
		n >>= 1
	}

	return p
}
