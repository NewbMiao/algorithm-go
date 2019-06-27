package leetcode

import (
	"math"
)

//https://leetcode-cn.com/problems/is-subsequence/
/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
*/
//经典dp 内存占用多
func isSubsequence(s string, t string) bool {
	row := len(t)
	col := len(s)
	if col == 0 {
		return true
	}
	if row == 0 {
		return false
	}
	dp := make([][]int, row)
	for k := range dp {
		dp[k] = make([]int, col)
	}

	if s[0] == t[0] {
		dp[0][0] = 1
	}
	for i := 1; i < row; i++ {
		if s[0] == t[i] {
			dp[i][0] = 1
		}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i][0])))
	}

	for i := 1; i < col; i++ {
		if s[i] == t[0] {
			dp[0][i] = 1
		}
		dp[0][i] = int(math.Max(float64(dp[0][i-1]), float64(dp[0][i])))
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			if s[j] == t[i] {
				dp[i][j] = int(math.Max(float64(dp[i-1][j-1]+1), float64(dp[i][j])))
			}
		}
	}
	if dp[row-1][col-1] == col {
		return true
	}
	return false
}

//压缩dp 减少内存占用
func isSubsequence2(s string, t string) bool {
	row := len(t)
	col := len(s)
	if col == 0 {
		return true
	}
	if row == 0 {
		return false
	}
	dp := make([]int, col)

	if s[0] == t[0] {
		dp[0] = 1
	}
	for i := 1; i < col; i++ {
		if t[0] == s[i] {
			dp[i] = 1
		}
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i])))
	}

	pre := 0
	for i := 1; i < row; i++ {
		pre = dp[0]
		if t[i] == s[0] {
			dp[0] = 1
		}

		for j := 1; j < col; j++ {
			tmp := dp[j]
			dp[j] = int(math.Max(float64(dp[j]), float64(dp[j-1])))
			if s[j] == t[i] {
				dp[j] = int(math.Max(float64(pre+1), float64(dp[j])))
			}
			pre = tmp
			if dp[j] == col {
				return true
			}
		}
	}
	if dp[col-1] == col {
		return true
	}
	return false
}

//双指针遍历 更快
func isSubsequence3(s string, t string) bool {
	var si, ti int
	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si += 1
			ti += 1
		} else {
			ti += 1
		}
	}

	if si == len(s) {
		return true
	}

	return false
}
