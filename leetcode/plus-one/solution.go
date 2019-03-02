package solution

func plusOne(digits []int) []int {
	incr := append([]int{0}, digits...)
	carry := 1
	for i := len(digits); i >= 0; i-- {
		s := incr[i] + carry
		incr[i] = s % 10
		carry = s / 10
	}
	if incr[0] == 0 {
		incr = incr[1:]
	}
	return incr
}
