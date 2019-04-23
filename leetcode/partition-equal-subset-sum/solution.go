package solution

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	// Find total sum.
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 == 1 {
		return false
	}
	target := total / 2

	// Produce set of all subsums.
	subsums := map[int]bool{0: true}
	for _, num := range nums {
		next := map[int]bool{}
		for subsum := range subsums {
			next[subsum] = true
			next[subsum+num] = true
		}
		subsums = next

		if subsums[target] {
			return true
		}
	}

	return false
}
