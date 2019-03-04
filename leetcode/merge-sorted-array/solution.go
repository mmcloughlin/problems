package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	out := nums1
	i1, i2, i := m-1, n-1, m+n-1

	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] < nums2[i2] {
			out[i] = nums2[i2]
			i2--
		} else {
			out[i] = nums1[i1]
			i1--
		}
		i--
	}

	if i2 >= 0 {
		copy(out, nums2[:i2+1])
	}
}
