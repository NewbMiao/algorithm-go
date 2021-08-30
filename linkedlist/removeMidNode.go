package linkedlist

import (
	"math"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

// 删除列表中间节点.
func RemoveMidNode(l *linkedlist.LList) {
	// 空节点或单节点不移除
	if l == nil || l.Head == nil || l.Head.Next == nil {
		return
	}

	// 双节点移除头结点
	if l.Head.Next.Next == nil {
		l.Head = l.Head.Next
		return
	}

	// 每增加2节点，删除节点后移一个
	pre := l.Head
	cur := l.Head.Next.Next
	for cur.Next != nil && cur.Next.Next != nil {
		pre = pre.Next
		cur = cur.Next.Next
	}
	pre.Next = pre.Next.Next
}

// 按比例删除节点.
func RemoveByRatio(l *linkedlist.LList, a, b int) {
	n := 0 // n=l.Size
	cur := l.Head
	for cur != nil {
		n++
		cur = cur.Next
	}

	// 要删除节点
	r := math.Ceil(float64(a*n) / float64(b))
	if r == 1 {
		l.Head = l.Head.Next
		return
	}
	cur = l.Head
	for {
		r--
		if r == 1 {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
}
