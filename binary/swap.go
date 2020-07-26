package binary

//交换二进制指定区间 (i+n-1~i) 和 （j+n-1~j）.
func swapBitSeq(i, j, n, b uint) (r uint) {
	x := (b>>i ^ b>>j) & (1<<n - 1)
	r = b ^ (x<<i | x<<j)
	return
}
