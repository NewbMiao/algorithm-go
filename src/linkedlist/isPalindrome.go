package linkedlist

import (
	"kit/stack"
	. "kit/linkedlist"
)

//链表是否为回文结构 栈
func isPalindrome1(l *LList) bool {
	if l == nil || l.IsEmpty() {
		return false
	}
	st := stack.New()
	cur := l.Head
	for cur != nil {
		st.Push(cur.Value)
		cur = cur.Next
	}
	cur = l.Head
	for cur != nil {
		tmp, _ := st.Pop()
		if cur.Value != tmp {
			return false
		}
		cur = cur.Next
	}
	return true
}

//栈压入一半
func isPalindrome2(l *LList) bool {
	if l == nil || l.IsEmpty() {
		return false
	}
	n := l.Size
	m := n / 2
	st := stack.New()
	cur := l.Head
	for cur != nil {
		if m > 0 {
			st.Push(cur.Value)
		} else if m == 0 {
			if n%2 == 1 { //跳过中点
				cur = cur.Next
				continue
			}
		} else {
			top, _ := st.Pop()
			if top != cur.Value {
				return false
			}
		}
		m--
		cur = cur.Next
	}
	return true
}
