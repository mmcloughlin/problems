package prod

func productExceptSelf(nums []int) []int {
	n := len(nums)

	// Output array will be used for intermediate storage.
	output := make([]int, len(nums))

	// Compute tail products.
	r := 1
	for i := n - 1; i >= 0; i-- {
		output[i] = r
		r *= nums[i]
	}

	// Now compute running sums from the left.
	l := 1
	for i := 0; i < n; i++ {
		output[i] *= l
		l *= nums[i]
	}

	return output
}
