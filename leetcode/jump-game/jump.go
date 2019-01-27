package jump

func canJump(nums []int) bool {
	d := 0 // min required distance
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] >= d {
			d = 0
		}
		d++
	}
	return nums[0] >= d
}
