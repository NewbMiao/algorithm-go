package linkedlist

import (
	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
	"github.com/NewbMiao/algorithm-go/kit/stack"
)

//链表是否为回文结构 栈
func isPalindrome1(l *linkedlist.LList) bool {
	if l == nil || l.IsEmpty() || l.Head.Next == nil {
		return true
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
func isPalindrome2_1(l *linkedlist.LList) bool {
	if l == nil || l.IsEmpty() || l.Head.Next == nil {
		return true
	}
	n := l.Size
	m := n / 2

	st := stack.New()
	cur := l.Head
	for cur != nil {
		if m > 0 {
			st.Push(cur.Value)
		} else if m == 0 && n%2 == 1 { //跳过中点

		} else if m <= 0 {
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

//栈压入一半
func isPalindrome2_2(l *linkedlist.LList) bool {
	if l == nil || l.IsEmpty() || l.Head.Next == nil {
		return true
	}

	st := stack.New()
	cur := l.Head
	right := l.Head.Next //右半区
	for cur.Next != nil && cur.Next.Next != nil {
		cur = cur.Next.Next
		right = right.Next
	}

	for right != nil {
		st.Push(right.Value)
		right = right.Next
	}
	cur = l.Head
	for !st.IsEmpty() {
		top, _ := st.Pop()
		if top != cur.Value {
			return false
		}
		cur = cur.Next
	}
	return true
}

//逆置右半区后比较再恢复
func isPalindrome3(l *linkedlist.LList) bool {
	if l == nil || l.IsEmpty() || l.Head.Next == nil {
		return true
	}
	n1 := l.Head
	n2 := l.Head

	for n2.Next != nil && n2.Next.Next != nil {
		n2 = n2.Next.Next //结尾
		n1 = n1.Next      //中点
	}
	n2 = n1.Next //右边第一个节点
	n1.Next = nil

	//逆置右半边
	var n3 *linkedlist.LNode
	for n2 != nil {
		n3 = n2.Next
		n2.Next = n1
		n1 = n2 //n1移动
		n2 = n3 //n2移动
	}

	n3 = n1     //右边最后节点
	n2 = l.Head //左边头结点
	res := true
	for n1 != nil && n2 != nil {
		if n1.Value != n2.Value {
			res = false
			break
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	//恢复逆置
	n1 = nil
	n2 = nil
	for n3 != nil {
		n1 = n3.Next
		n3.Next = n2
		n2 = n3
		n3 = n1
	}

	return res
}
