package binary

const intSize = 32 << (^uint(0) >> 63)

func ReverseBin32(v uint32) (r uint32) {
	var s uint32 = 32
	var mask uint32 = ^uint32(0)
	// fmt.Printf("%b,%d,%d\n", mask, mask, ^0)
	// 11111111111111111111111111111111,4294967295,-1
	s >>= 1
	for s > 0 {
		mask ^= (mask << s)
		v = ((v >> s) & mask) | ((v << s) & (^mask));
		s >>= 1
	}

	return v
}
