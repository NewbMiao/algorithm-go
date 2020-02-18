package linkedlist

import (
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/linkedlist"
)

//原地逆置链表
func ReverseList(l *LList) {
	if l == nil || l.Head == nil {
		return
	}
	head := l.Head
	var pre, next *LNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	l.Head = pre
}

//原地逆置双链表
func ReverseDList(l *DList) {
	if l == nil || l.Head == nil {
		return
	}
	head := l.Head
	var pre, next *DoubleNode
	for head != nil {
		next = head.Next
		head.Next = pre
		head.Last = next
		pre = head
		head = next
	}
	tail := l.Head
	l.Head = pre
	//头结点前指针指向尾节点
	l.Head.Last = tail
}
