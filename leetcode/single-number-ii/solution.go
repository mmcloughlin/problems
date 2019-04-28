package solution

func singleNumber(nums []int) int {
	x := uint64(0)
	for b := uint(0); b < 64; b++ {
		var c uint64
		for _, num := range nums {
			c += (uint64(num) >> b) & 1
		}
		x |= (c % 3) << b
	}
	return int(x)
}
