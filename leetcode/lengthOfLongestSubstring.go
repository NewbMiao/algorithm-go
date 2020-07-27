package leetcode

import "math"

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。.
*/
func lengthOfLongestSubstring(s string) int {
	st := map[byte]int{} //s中字符索引map
	var i int            //重复字符索引较大者
	var ans int          //不重复字符长度
	for j := 0; j < len(s); j++ {
		if si, ok := st[s[j]]; ok {
			i = int(math.Max(float64(si), float64(i)))
		}
		ans = int(math.Max(float64(ans), float64(j-i+1)))
		st[s[j]] = j + 1 //+1: 防止s长度为1时无法正确计算
	}
	return ans
}
