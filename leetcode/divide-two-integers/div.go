package div

import "math"

func divide(dividend int, divisor int) int {
	// Determine sign, and ensure operands are positive.
	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}

	// Find least shift of divisor that exceeds dividend.
	b := 1
	s := divisor
	for s < dividend {
		b <<= 1
		s <<= 1
	}

	// Successively remove shifts of divisor.
	result := 0
	rem := dividend
	for rem >= divisor {
		for s > rem {
			s >>= 1
			b >>= 1
		}
		result |= b
		rem -= s
	}

	// Apply sign.
	if sign < 0 {
		result = -result
	}

	// Check for overflow.
	if result < math.MinInt32 || result > math.MaxInt32 {
		return math.MaxInt32
	}

	return result
}
