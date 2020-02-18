package string_alg

func GetIndexOf(ss, ms string) (i int) {
	if ss == "" || ms == "" || len(ms) > len(ss) {
		return -1
	}
	var si, mi int
	next := getNextArr(ms)
	for si < len(ss) && mi < len(ms) {
		if ss[si] == ms[mi] {
			si += 1
			mi += 1
		} else if next[mi] == -1 { //无公共匹配串，str后移一个字符
			si++
		} else { //有公共子串，回退到上一个公共子串位置
			mi = next[mi]
		}
	}
	if mi == len(ms) {
		return si - mi
	}
	return -1
}

//match[i]之前的前缀和后缀最大匹配公共长度
func getNextArr(ms string) (next []int) {
	mlen := len(ms)
	if mlen == 0 {
		return []int{-1}
	}
	next = make([]int, mlen)
	next[0] = -1
	next[1] = 0
	pos := 2
	cn := 0
	for pos < mlen {
		if ms[pos-1] == ms[cn] { //前后缀匹配，则前移一位
			cn += 1
			next[pos] = cn
			pos++
		} else if cn > 0 { //不匹配，则回退到上一次匹配点cn
			cn = next[cn]
		} else { //回退完也不匹配
			next[pos] = 0
			pos++
		}
	}
	return
}
