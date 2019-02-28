package solution

func trailingZeroes(n int) int {
	tz := 0
	d := 5
	for d <= n {
		tz += n / d
		d *= 5
	}
	return tz
}
