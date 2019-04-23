package solution

func findTargetSumWays(nums []int, S int) int {
	counts := map[int]int{0: 1}
	for _, num := range nums {
		next := map[int]int{}
		for s, c := range counts {
			next[s+num] += c
			next[s-num] += c
		}
		counts = next
	}
	return counts[S]
}
