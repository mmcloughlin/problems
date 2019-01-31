package gas

func canCompleteCircuit(gas []int, cost []int) int {
	mn := 0
	roundtrip := 0
	for i := range gas {
		roundtrip += gas[i] - cost[i]
		mn = min(roundtrip, mn)
	}

	for i := range gas {
		if mn >= 0 {
			return i
		}
		mn = min(roundtrip, mn-gas[i]+cost[i])
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
