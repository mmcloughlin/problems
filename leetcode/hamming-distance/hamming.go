package hamming

func hammingDistance(x, y int) int {
	return popcnt(uint64(x ^ y))
}

func popcnt(x uint64) int {
	x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x & 0x0f0f0f0f0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f0f0f0f0f)
	x = (x & 0x00ff00ff00ff00ff) + ((x >> 8) & 0x00ff00ff00ff00ff)
	x = (x & 0x0000ffff0000ffff) + ((x >> 16) & 0x0000ffff0000ffff)
	x = (x & 0x00000000ffffffff) + ((x >> 32) & 0x00000000ffffffff)
	return int(x)
}
