package crc

const POLY  = 0x4C11DB7
func bitRev(d, bw uint) (r uint) {
	var i uint
	for ; i < bw; i++ {
		if d&1 == 1 {
			r |= 1 << (bw - i - 1)
		}
		d >>= 1
	}
	return
}

func makeCrcTbl() [256]uint {
	var i,j int
	var s uin
	var s [256]uint=[256]uint{0}
	var poly = bitRev(POLY,32)
	for i=0;i<256;i++{
		for j=0;i<8;j++{
			if i&1==1{

			}
		}
	}
	return
}