package perms

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	perms := [][]int{}
	for i := range nums {
		nums[0], nums[i] = nums[i], nums[0]
		for _, perm := range permute(nums[1:]) {
			perms = append(perms, append([]int{nums[0]}, perm...))
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return perms
}
