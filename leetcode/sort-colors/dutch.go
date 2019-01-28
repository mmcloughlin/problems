package dutch

func sortColors(nums []int) {
	n0, n01 := 0, 0 // num 0, num 0 and 1

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v < 2 {
			nums[i] = 2
			nums[n01] = 1
			n01++
		}
		if v < 1 {
			nums[n0] = 0
			n0++
		}
	}
}
