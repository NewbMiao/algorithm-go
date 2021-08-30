package binary

const intSize = 32 << (^uint(0) >> 63)

func reverseBit1(x uint) (r uint) {
	r = x
	var s uint = 32 << (^uint(0) >> 63)
	s--
	for x >>= 1; x != 0; x >>= 1 {
		r <<= 1
		r |= x & 1
		s--
	}
	r <<= s // shift when v's highest bits are zero
	return
}

// look reverse table
// https://stackoverflow.com/questions/15107398/algorithm-behind-the-generation-of-the-reverse-bits-lookup-table8-bit
func reverseBit2(x uint) (r uint) {
	// reverse map
	rMap := make([]uint, 0)
	p2 := func(n uint) []uint {
		return []uint{n, n + 2*64, n + 1*64, n + 3*64}
	}
	p4 := func(n uint) (r []uint) {
		r = append(r, p2(n)...)
		r = append(r, p2(n+2*16)...)
		r = append(r, p2(n+1*16)...)
		r = append(r, p2(n+3*16)...)
		return r
	}
	p6 := func(n uint) (r []uint) {
		r = append(r, p4(n)...)
		r = append(r, p4(n+2*4)...)
		r = append(r, p4(n+1*4)...)
		r = append(r, p4(n+3*4)...)
		return r
	}
	rMap = append(rMap, p6(0)...)
	rMap = append(rMap, p6(2)...)
	rMap = append(rMap, p6(1)...)
	rMap = append(rMap, p6(3)...)

	var intSize uint = 32 << (^uint(0) >> 63)
	intSize -= 8
	for x != 0 {
		r |= rMap[x&0xff] << intSize
		x >>= 8
	}
	return r
}

func ReverseBin32(v uint32) (r uint32) {
	var s uint32 = 32
	mask := ^uint32(0)
	// fmt.Printf("%b,%d,%d\n", mask, mask, ^0)
	// 11111111111111111111111111111111,4294967295,-1
	s >>= 1
	for s > 0 {
		mask ^= (mask << s)
		v = ((v >> s) & mask) | ((v << s) & (^mask))
		s >>= 1
	}

	return v
}
