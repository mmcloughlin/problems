package missing

func findDisappearedNumbers(x []int) []int {
	const (
		notseen = 0
		seen    = -1
	)
	for i := 0; i < len(x); i++ {
		if x[i] < 1 {
			continue
		}
		j := x[i] - 1
		x[i] = notseen
		for j >= 0 {
			j, x[j] = x[j]-1, seen
		}
	}
	d := []int{}
	for i, v := range x {
		if v == notseen {
			d = append(d, i+1)
		}
	}
	return d
}
