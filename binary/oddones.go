package binary

import "fmt"

// http://graphics.stanford.edu/~seander/bithacks.html

//平行计算,异或压缩奇偶信息，压缩顺序:1->2->4->8->16->32.
func judgeBinOddOne1(x uint) bool {
	const intSize = 32 << (^uint(0) >> 63)
	var i uint = 1
	for i < intSize {
		x ^= x >> i
		i *= 2
	}
	return x&1 == 1
}

// 一个二进制数减1，若右起第一个1的索引位置是m，那么从m到0位的二进制会逐位取反。
// 所以，x &= x-1会使得从m到0位的二进制都会变成0。.
func judgeBinOddOne2(x uint) bool {
	cnt := 0
	for x > 0 {
		cnt++
		x &= x - 1
	}
	fmt.Printf("judgeBinOddOne2 has one: %d\n", cnt)
	return cnt&1 == 1
}

//查表法：256区间内1的个数表.
func judgeBinOddOne3(x uint) bool {
	var bitmap256 = [256]int{0}
	for i := 1; i < 256; i++ {
		bitmap256[i] = i&1 + bitmap256[i/2]
	}
	cnt := 0
	for x != 0 {
		cnt += bitmap256[x&0xff]
		x >>= 8
	}
	fmt.Printf("judgeBinOddOne has one: %d\n", cnt)
	return cnt&1 == 1
}

// 压缩相加
// 32位为例，64进制类似
// 1: 0x55555555 = 01010101 01010101 01010101 01010101  （循环序列01）
// 2: 0x33333333 = 00110011 00110011 00110011 00110011	（循环序列0011）
// 4: 0x0F0F0F0F = 00001111 00001111 00001111 00001111	（循环序列00001111）
// 8: 0x00FF00FF = 00000000 11111111 00000000 11111111	（...）
//16: 0x0000FFFF = 00000000 00000000 11111111 11111111	（...）
//
// (n&0x55555555)+((n>>1)&0x55555555)
// 将32位数中的32个段从左往右把相邻的两个段的值相加后放在2bits中，
// 就变成了16个段，每段2位。
// 同理(n&0x33333333)+((n>>2)&0x33333333)将16个段中相邻的两个段两两相加，
// 存放在4bits中，就变成了8个段，每段4位。
// 以此类推，最终求得数中1的个数就存放在一个段中，这个段32bits，.
func judgeBinOddOne4(x uint) bool {
	const intSize = 32 << (^uint(0) >> 63)
	var bMap = map[int][]uint{
		32: {0x55555555, 0x33333333, 0x0F0F0F0F, 0x00FF00FF, 0x0000FFFF},
		64: {0x5555555555555555, 0x3333333333333333, 0x0F0F0F0F0F0F0F0F, 0x00FF00FF00FF00FF, 0x0000FFFF0000FFFF, 0x00000000FFFFFFFF},
	}

	b := bMap[intSize]
	for offset, mask := range b {
		x = x&mask + x>>(1<<uint(offset))&mask
	}
	fmt.Printf("judgeBinOddOne4 has one: %d\n", x)

	return x&1 == 1
}

//parity table
//static const bool ParityTable256[256] =
//{
//#   define P2(n) n, n^1, n^1, n
//#   define P4(n) P2(n), P2(n^1), P2(n^1), P2(n)
//#   define P6(n) P4(n), P4(n^1), P4(n^1), P4(n)
//	  P6(0), P6(1), P6(1), P6(0)
//};
// https://stackoverflow.com/questions/53975983/parity-lookup-table-generation
//4bit parity generator: https://www.youtube.com/watch?v=RfTGvpY2Z5Y

func judgeBinOddOne5(x uint) bool {
	//parity map
	var pMap = make([]uint, 0)
	var p2 = func(n uint) []uint {
		return []uint{n, n ^ 1, n ^ 1, n}
	}
	var p4 = func(n uint) (r []uint) {
		r = append(r, p2(n)...)
		r = append(r, p2(n^1)...)
		r = append(r, p2(n^1)...)
		r = append(r, p2(n)...)
		return r
	}
	var p6 = func(n uint) (r []uint) {
		r = append(r, p4(n)...)
		r = append(r, p4(n^1)...)
		r = append(r, p4(n^1)...)
		r = append(r, p4(n)...)
		return r
	}
	pMap = append(pMap, p6(0)...)
	pMap = append(pMap, p6(1)...)
	pMap = append(pMap, p6(1)...)
	pMap = append(pMap, p6(0)...)

	//OR, for 32-bit words:
	//unsigned int v;
	//v ^= v >> 16;
	//v ^= v >> 8;
	//bool parity = ParityTable256[v & 0xff];

	//unsigned char * p = (unsigned char *) &v;
	//parity = ParityTable256[p[0] ^ p[1] ^ p[2] ^ p[3]];
	var p uint
	for x != 0 {
		p ^= pMap[x&0xff]
		x >>= 8
	}
	return p == 1
}

//并行查奇偶表：0110 1001 1001 0110 (0x6996 in hex).
func judgeBinOddOne6(x uint) bool {
	const intSize = 32 << (^uint(0) >> 63)
	if intSize == 64 {
		x ^= x >> 32
	}
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x &= 0xf
	return (0x6996>>x)&1 == 1
}

//multiple 8 operations.
func judgeBinOddOne7(x uint) bool {
	const intSize = 32 << (^uint(0) >> 63)
	iMap := map[int]uint{
		32: 0x11111111,
		64: 0x1111111111111111,
	}
	x ^= x >> 1
	x ^= x >> 2
	x = x & iMap[intSize] * iMap[intSize]
	return x>>(intSize-4)&1 == 1
}
