package leetcode

func reverseString(s []byte) {
	l := len(s)
	if l <= 1 {
		return
	}

	left, right := 0, l-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
