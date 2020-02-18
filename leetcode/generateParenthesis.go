package leetcode

import (
	"fmt"
	"sort"
)

/*
https://leetcode-cn.com/problems/generate-parentheses/

给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
*/
//回溯法
func generateParenthesis(n int) []string {

	ans := []string{}
	backtrackStr(&ans, "", 0, 0, n)
	return ans
}

func backtrackStr(ans *[]string, cur string, open, close, max int) {
	if len(cur) == 2*max {
		*ans = append(*ans, cur)
		return
	}
	if open < max {
		backtrackStr(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		backtrackStr(ans, cur+")", open, close+1, max)
	}
}

//闭合数
func generateParenthesis1(n int) []string {
	ans := []string{}
	if n == 0 {
		return []string{""}
	}
	for c := 0; c < n; c++ {
		for _, left := range generateParenthesis1(c) {
			for _, right := range generateParenthesis1(n - c - 1) {
				ans = append(ans, fmt.Sprintf("(%s)%s", left, right))
			}
		}
	}
	sort.Strings(ans)
	return ans
}
