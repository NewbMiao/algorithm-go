package linkedlist

import "fmt"

type DoubleNode struct {
	Value      interface{}
	Last, Next *DoubleNode
}

type DList struct {
	Head *DoubleNode
	Size int
}

func NewDList() *DList {
	return &DList{}
}

func NewDNode(n interface{}) *DoubleNode {
	return &DoubleNode{Value: n}
}
func (l *DList) PushBack(n interface{}) {
	if n == nil {
		return
	}
	if l.Head == nil {
		l.Head = NewDNode(n)
		l.Head.Last = l.Head
	} else {
		pre := l.Head.Last
		node := NewDNode(n)
		pre.Next = node
		node.Last = pre
		l.Head.Last = node
	}
	l.Size += 1
}

func (l *DList) PushFront(n interface{}) {
	if n == nil {
		return
	}
	if l.Head == nil {
		l.Head = NewDNode(n)
		l.Head.Last = l.Head
	} else {
		next := l.Head
		node := NewDNode(n)
		node.Next = next
		node.Last = next.Last
		next.Last = node

		l.Head = node

	}
	l.Size += 1
}
func (l *DList) PopBack() (res interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	node := l.Head.Last
	res = node.Value

	if node == node.Last {
		l.Head = nil
	} else {
		l.Head.Last = node.Last
		l.Head.Last.Next = nil
	}

	l.Size -= 1
	return
}
func (l *DList) PopFront() (res interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	res = l.Head.Value

	node := l.Head
	if node == node.Last {
		l.Head = nil
	} else {
		l.Head = node.Next
		l.Head.Last = node.Last
	}
	l.Size -= 1
	return
}

func (l *DList) Front() *DoubleNode {
	return l.Head
}

func (l *DList) Back() *DoubleNode {
	if l.Head == nil {
		return nil
	}
	return l.Head.Last
}

func (l *DList) IsEmpty() bool {
	return l.Head == nil
}

func (l *DList) IsLast(n *DoubleNode) bool {
	if l.Head == nil {
		return false
	}
	return l.Head.Last == n
}

func (l *DList) String() string {
	tmp := l.Head
	r := []interface{}{}
	for tmp != nil {
		r = append(r, tmp.Value)
		tmp = tmp.Next
	}
	return fmt.Sprintf("%v", r)
}
