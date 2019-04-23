package solution

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		if target == 0 {
			return [][]int{{}}
		}
		return nil
	}

	top := len(candidates) - 1
	n := candidates[top]
	others := candidates[:top]
	chosen := []int{}

	var combinations [][]int

	for target >= 0 {
		for _, sub := range combinationSum(others, target) {
			combination := append([]int{}, sub...)
			combination = append(combination, chosen...)
			combinations = append(combinations, combination)
		}

		chosen = append(chosen, n)
		target -= n
	}

	return combinations
}
