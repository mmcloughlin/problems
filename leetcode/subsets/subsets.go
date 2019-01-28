package subsets

func subsets(nums []int) [][]int {
	powerset := [][]int{{}}
	for _, num := range nums {
		next := make([][]int, 0, len(powerset))
		for _, set := range powerset {
			newset := append([]int{}, set...)
			newset = append(newset, num)
			next = append(next, newset)
		}
		powerset = append(powerset, next...)
	}
	return powerset
}
