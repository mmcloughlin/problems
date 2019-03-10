package solution

func checkSubarraySum(nums []int, k int) bool {
	if k == 0 {
		for i := 1; i < len(nums); i++ {
			if nums[i-1] == 0 && nums[i] == 0 {
				return true
			}
		}
		return false
	}

	mod := map[int]bool{
		nums[0] % k: true,
	}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		nxt := map[int]bool{}

		for m := range mod {
			nxt[(m+num)%k] = true
		}

		if nxt[0] {
			return true
		}

		nxt[num%k] = true
		mod = nxt
	}

	return false
}
