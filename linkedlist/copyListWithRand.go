package linkedlist

import (
	"fmt"
)

type randNode struct {
	Next  *randNode
	Value int
	Rand  *randNode
}

// 拷贝带有rand的节点.
func copyListWithRand1(head *randNode) *randNode {
	if head == nil {
		return nil
	}
	m := make(map[*randNode]*randNode)
	cur := head
	for cur != nil {
		m[cur] = &randNode{Value: cur.Value}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		m[cur].Rand = m[cur.Rand]
		m[cur].Next = m[cur.Next]
		cur = cur.Next
	}
	return m[head]
}

func copyListWithRand2(head *randNode) *randNode {
	if head == nil {
		return nil
	}

	// 复制节点带每一个节点后
	cur := head
	var next *randNode
	for cur != nil {
		next = cur.Next
		cur.Next = &randNode{
			Next:  next,
			Value: cur.Value,
			Rand:  nil,
		}
		cur = next
	}

	// copy rand
	cur = head
	for cur != nil {
		if cur.Rand != nil {
			cur.Next.Rand = cur.Rand.Next
		}
		cur = cur.Next.Next
	}

	// 分离
	cur = head
	res := cur.Next
	var curCopy *randNode
	for cur != nil {
		curCopy = cur.Next
		cur.Next = curCopy.Next
		if curCopy.Next == nil {
			curCopy.Next = nil
		} else {
			curCopy.Next = curCopy.Next.Next
		}
		cur = cur.Next
	}
	return res
}

// 按[][cur,next,rand]方式打印.
func (s *randNode) String() string {
	cur := s
	res := [][3]int{}
	for cur != nil {
		tmp := [3]int{cur.Value}
		if cur.Next != nil {
			tmp[1] = cur.Next.Value
		}
		if cur.Rand != nil {
			tmp[2] = cur.Rand.Value
		}

		res = append(res, tmp)
		cur = cur.Next
	}
	return fmt.Sprint(res)
}
