package solution

func getSum(a int, b int) int {
	s := uint64(0)
	A := uint64(a)
	B := uint64(b)
	C := uint64(0)
	i := uint(0)
	for A != 0 || B != 0 || C != 0 {
		bit := (A ^ B ^ C) & 1
		s |= bit << i

		C = maj(A, B, C) & 1
		i++
		A >>= 1
		B >>= 1
	}
	return int(s)
}

func maj(x, y, z uint64) uint64 {
	return (x & y) | (x & z) | (y & z)
}
