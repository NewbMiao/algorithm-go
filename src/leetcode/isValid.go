package leetcode

import "kit/stack"

/*
https://leetcode-cn.com/problems/valid-parentheses/
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

*/
func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	if l < 2 {
		return false
	}
	pMap := map[uint8]uint8{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	st := stack.New()

	var pre uint8
	for i := 0; i < l; i++ {

		if i > 0 && pMap[pre] == s[i] {
			st.Pop()
			pre, _ = st.Peek().(uint8)
		} else {
			st.Push(s[i])
			pre = s[i]
		}
	}
	if st.IsEmpty() {
		return true
	}
	return false
}
