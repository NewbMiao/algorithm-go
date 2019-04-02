package binary


//x & ~0 = x;	x & 0 = 0
func min(x, y int) (r int) {
	var less int
	if x < y {
		less = 1
	}
	r = y ^ ((x ^ y) & -less)
	return
}
func max(x, y int) (r int) {
	var less int
	if x < y {
		less = 1
	}
	r = x ^ ((x ^ y) & -less)
	return
}

func isPowerOf2(x int) bool {
	return x != 0 && x&(x-1) == 0
}


