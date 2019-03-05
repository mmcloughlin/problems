package solution

func rotate(nums []int, k int) {
	n := len(nums)
	m := 0
	i := 0
	for m < n {
		p := nums[i]
		for j := (i + k) % n; ; j = (j + k) % n {
			nums[j], p = p, nums[j]
			m++
			if j == i {
				break
			}
		}
		i++
	}
}
