package crc

import "fmt"

const POLY uint32 = 0x4C11DB7 //  birrev:0xedb88320 (IEEE)

func bitRev(d uint32, bw uint) (r uint32) {
	var i uint
	for ; i < bw; i++ {
		if d&1 == 1 {
			r |= 1 << (bw - i - 1)
		}
		d >>= 1
	}
	return
}

func makeCrcTbl() [256]uint32 {
	var i, j uint32
	var crc uint32
	var s = [256]uint32{0}
	var poly = bitRev(POLY, 32)
	fmt.Printf("poly 0x%x bitrev is 0x%x,\n", POLY, poly)
	for i = 0; i < 256; i++ {
		crc = i
		for j = 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		s[i] = crc
	}

	return s
}

func Crc32(data []byte) uint32 {
	var crc uint32 = ^uint32(0) //0xFFFFFFFF
	table := makeCrcTbl()
	for i, cnt := 0, len(data); i < cnt; i++ {
		crc = (crc >> 8) ^ table[data[i]^byte(crc)]
	}
	return ^crc //反转 等效 crc ^ 0xFFFFFFFF
}
