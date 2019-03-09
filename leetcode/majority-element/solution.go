package solution

func majorityElement(nums []int) int {
	n := len(nums)
	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
		if counts[num] > n/2 {
			return num
		}
	}
	panic("unreachable")
}
