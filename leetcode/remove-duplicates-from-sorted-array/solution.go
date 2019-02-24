package solution

func removeDuplicates(nums []int) int {
	i := 0
	for _, num := range nums {
		if num != nums[i] {
			i++
			nums[i] = num
		}
	}
	return i + 1
}
