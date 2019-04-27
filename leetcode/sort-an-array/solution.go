package solution

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	n := len(nums)
	m := n / 2
	a := sortArray(nums[:m])
	b := sortArray(nums[m:])

	s := make([]int, 0, n)
	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			s = append(s, a[0])
			a = a[1:]
		} else {
			s = append(s, b[0])
			b = b[1:]
		}
	}

	if len(a) > 0 {
		s = append(s, a...)
	}
	if len(b) > 0 {
		s = append(s, b...)
	}

	return s
}
