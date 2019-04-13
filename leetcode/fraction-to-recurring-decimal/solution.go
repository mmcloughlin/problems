package solution

type rational struct{ num, den int }

func (r *rational) mul(m int) {
	r.num *= m
}

func (r *rational) reduce() int {
	w := r.num / r.den
	r.num %= r.den
	return w
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	// Handle sign.
	sign := 1
	if numerator < 0 {
		sign *= -1
		numerator = -numerator
	}
	if denominator < 0 {
		sign *= -1
		denominator = -denominator
	}

	s := ""
	if sign < 0 {
		s += "-"
	}

	// Handle whole part.
	f := rational{num: numerator, den: denominator}
	w := f.reduce()

	s += numeral(w)
	if f.num == 0 {
		return s
	}

	// Fractional part.
	frac := ""
	seen := map[rational]int{}
	i := 0
	for f.num != 0 {
		if j, found := seen[f]; found {
			frac = frac[:j] + "(" + frac[j:] + ")"
			break
		}
		seen[f] = i
		f.mul(10)
		d := f.reduce()
		frac += string([]byte{digit(d)})
		i++
	}

	return s + "." + frac
}

func numeral(x int) string {
	s := ""
	for {
		s = string([]byte{digit(x % 10)}) + s
		x /= 10
		if x == 0 {
			break
		}
	}
	return s
}

func digit(d int) byte {
	return '0' + byte(d)
}
