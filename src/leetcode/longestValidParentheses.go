package leetcode

import (
	"kit/stack"
	"math"
)

/*
https://leetcode-cn.com/problems/longest-valid-parentheses/submissions/
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

https://leetcode-cn.com/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode/
*/
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	st := stack.New()
	max := 0
	st.Push(-1)
	for i, v := range s {
		if v == '(' {
			st.Push(i)
		} else {
			st.Pop()
			if st.IsEmpty() {
				st.Push(i)
			} else {
				max = int(math.Max(float64(max), float64(i-st.Peek().(int))))
			}
		}
	}
	return max
}

func longestValidParentheses1(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}
	max := 0
	dp := make([]int, l)
	for i := 1; i < l; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { //...()
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if s[i-1] == ')' { // ...))
				if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' { // (..))
					if i-dp[i-1] >= 2 { //dp[i-1]+2 是i位置的长度
						dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		max = int(math.Max(float64(max), float64(dp[i])))
	}
	return max
}

func longestValidParentheses2(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}
	max := 0
	var left, right int
	//从左到右
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			left += 1
		} else {
			right += 1
		}
		if left == right {
			max = int(math.Max(float64(max), float64(2*right)))
		} else if right > left {
			left = 0
			right = 0
		}
	}

	//从右到左
	left=0
	right=0
	for i := l - 1; i >= 0; i-- {
		if s[i] == '(' {
			left += 1
		} else {
			right += 1
		}
		if left == right {
			max = int(math.Max(float64(max), float64(2*right)))
		} else if right < left {
			left = 0
			right = 0
		}
	}
	return max
}
