package linkedlist

import (
	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

// 原地逆置链表.
func ReverseList(l *linkedlist.LList) {
	if l == nil || l.Head == nil {
		return
	}
	head := l.Head
	var pre, next *linkedlist.LNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	l.Head = pre
}

// 原地逆置双链表.
func ReverseDList(l *linkedlist.DList) {
	if l == nil || l.Head == nil {
		return
	}
	head := l.Head
	var pre, next *linkedlist.DoubleNode
	for head != nil {
		next = head.Next
		head.Next = pre
		head.Last = next
		pre = head
		head = next
	}
	tail := l.Head
	l.Head = pre
	// 头结点前指针指向尾节点
	l.Head.Last = tail
}
