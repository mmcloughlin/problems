package movz

func moveZeroes(x []int) {
	i := 0
	for _, v := range x {
		if v != 0 {
			x[i] = v
			i++
		}
	}
	for ; i < len(x); i++ {
		x[i] = 0
	}
}
