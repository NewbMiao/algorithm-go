package linkedlist

import "fmt"

type LNode struct {
	Value interface{}
	Next  *LNode
}

type LList struct {
	Head *LNode
	Size int
}

func NewList() *LList {
	return &LList{}
}

func NewLNode(n interface{}) *LNode {
	return &LNode{Value: n}
}
func (l *LList) Push(n interface{}) {
	if n == nil {
		return
	}
	if l.Head == nil {
		l.Head = NewLNode(n)
	} else {
		pre := l.Head
		cur := l.Head.Next
		for cur != nil {
			pre = cur
			cur = cur.Next
		}
		pre.Next = NewLNode(n)
	}
	l.Size += 1
}

func (l *LList) Pop() (res interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	res = l.Head.Value

	l.Head = l.Head.Next
	l.Size -= 1
	return
}

func (l *LList) IsEmpty() bool {
	return l == nil || l.Head == nil
}

func (l *LList) Last() *LNode {
	if l == nil || l.Head == nil {
		return nil
	}
	n := l.Head
	for n.Next != nil && n.Next != l.Head {
		n = n.Next
	}
	return n
}

func (l *LList) IsLast(n *LNode) bool {
	if l == nil || l.Head == nil {
		return false
	}
	last := l.Head
	for last != nil {
		if last.Next == nil || last.Next == l.Head {
			break
		}
		last = last.Next
	}
	return last == n
}

func (l *LList) String() string {
	tmp := l.Head
	r := []interface{}{}
	for tmp != nil {
		r = append(r, tmp.Value)
		if tmp.Next == l.Head { //避免成环死循环
			break
		}
		tmp = tmp.Next
	}
	return fmt.Sprintf("%v", r)
}
