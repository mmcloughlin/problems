package solution

func subarraySum(nums []int, k int) int {
	type sum struct{ val, count int }
	sums := []sum{}
	total := 0

	for _, num := range nums {
		seen := false
		for i := range sums {
			sums[i].val += num
			if sums[i].val == num {
				sums[i].count++
				seen = true
			}
		}
		if !seen {
			sums = append(sums, sum{val: num, count: 1})
		}
		for i := range sums {
			if sums[i].val == k {
				total += sums[i].count
			}
		}
	}

	return total
}
