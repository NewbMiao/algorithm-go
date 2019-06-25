package linkedlist

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
	return l.Head == nil
}

func (l *LList) IsLast() bool {
	return l.Head.Next == nil
}
