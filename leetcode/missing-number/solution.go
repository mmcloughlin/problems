package solution

func missingNumber(nums []int) int {
	// Sum the array.
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// Sum of 0, 1, 2, ..., n.
	all := len(nums) * (len(nums) + 1) / 2

	return all - sum
}
