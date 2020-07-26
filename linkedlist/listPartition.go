package linkedlist

import (
	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

//链表按pivot分为小中大三部分
func listPartition1(l *linkedlist.LList, pivot int) {
	if l == nil || l.IsEmpty() {
		return
	}
	cur := l.Head
	n := 0
	arr := make([]*linkedlist.LNode, 0)
	for cur != nil {
		arr = append(arr, cur)
		n++
		cur = cur.Next
	}

	small := -1
	big := n
	index := 0

	for index < big {
		if arr[index].Value.(int) < pivot {
			small++
			arr[index], arr[small] = arr[small], arr[index]
			index++
		} else if arr[index].Value.(int) > pivot {
			big--
			arr[index], arr[big] = arr[big], arr[index]
		} else {
			index++
		}
	}
	for i := 1; i < len(arr); i++ {
		arr[i-1].Next = arr[i]
	}
	arr[len(arr)-1].Next = nil
	l.Head = arr[0]
}

//分隔中保持原有顺序
func listPartition2(l *linkedlist.LList, pivot int) {
	if l == nil || l.IsEmpty() {
		return
	}
	cur := l.Head
	var sH, sT, eH, eT, bH, bT, next *linkedlist.LNode

	for cur != nil {
		next = cur.Next
		cur.Next = nil
		if cur.Value.(int) < pivot {
			if sH == nil {
				sH = cur
				sT = cur
			} else {
				sT.Next = cur
				sT = cur
			}
		} else if cur.Value.(int) == pivot {
			if eH == nil {
				eH = cur
				eT = cur
			} else {
				eT.Next = cur
				eT = cur
			}
		} else {
			if bH == nil {
				bH = cur
				bT = cur
			} else {
				bT.Next = cur
				bT = cur
			}
		}
		cur = next
	}

	//连接小和中
	if sT != nil {
		sT.Next = eH
		if eT == nil { //中为空
			eT = sT
		}
	}
	//连接大
	if eT != nil {
		eT.Next = bH
	}

	if sH != nil {
		l.Head = sH
	} else if eH != nil {
		l.Head = eH
	} else {
		l.Head = bH
	}
}
