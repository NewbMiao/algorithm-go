package leetcode

/*
https://leetcode-cn.com/problems/regular-expression-matching/
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。.
*/
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' { //* : 校验0个或多个
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}
	return firstMatch && isMatch(s[1:], p[1:])
}
func isMatch2(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	dp := make([][]bool, len(s)+1)
	for k := range dp {
		dp[k] = make([]bool, len(p)+1)
	}

	dp[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}
