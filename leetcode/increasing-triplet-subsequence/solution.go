package solution

func increasingTriplet(nums []int) bool {
	const unknown = -1
	i, j := 0, unknown
	for l := 1; l < len(nums); l++ {
		switch {
		case j != unknown && nums[l] > nums[j]:
			return true
		case nums[l] < nums[i]:
			i = l
		case nums[i] < nums[l] && (j == unknown || nums[l] < nums[j]):
			j = l
		}
	}
	return false
}
