package solution

func reverseBits(x uint32) uint32 {
	x = ((x & 0x0000ffff) << 16) | ((x >> 16) & 0x0000ffff)
	x = ((x & 0x00ff00ff) << 8) | ((x >> 8) & 0x00ff00ff)
	x = ((x & 0x0f0f0f0f) << 4) | ((x >> 4) & 0x0f0f0f0f)
	x = ((x & 0x33333333) << 2) | ((x >> 2) & 0x33333333)
	x = ((x & 0x55555555) << 1) | ((x >> 1) & 0x55555555)
	return x
}
