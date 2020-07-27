package linkedlist

import "github.com/NewbMiao/algorithm-go/kit/linkedlist"

//原地逆置部分链表
func ReversePart(l *linkedlist.LList, from, to int) {
	if l == nil {
		return
	}
	n := 0
	cur := l.Head
	var fPre, tPos *linkedlist.LNode //先获取逆置开始的【前一个节点】和结束的【后一个节点】
	for cur != nil {
		n++
		if from-1 == n {
			fPre = cur
		}
		if to+1 == n {
			tPos = cur
		}
		cur = cur.Next
	}
	if from > to || from < 1 || to > n {
		return
	}

	//开始逆置节点
	pre := l.Head
	if fPre != nil {
		pre = fPre.Next
	}

	cur = pre.Next
	pre.Next = tPos //逆置后pre为tPos前一节点

	var next *linkedlist.LNode
	for cur != tPos {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	if fPre == nil { //头结点反转
		l.Head = pre
	} else {
		fPre.Next = pre
	}
}
