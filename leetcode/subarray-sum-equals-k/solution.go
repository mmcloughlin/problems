package solution

func subarraySum(nums []int, k int) int {
	total := 0
	prefixcount := map[int]int{0: 1}
	prefix := 0
	for _, num := range nums {
		prefix += num
		total += prefixcount[prefix-k]
		prefixcount[prefix]++
	}
	return total
}
