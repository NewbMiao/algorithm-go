package linkedlist

import . "kit/linkedlist"

func removeRepetition1(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	m := make(map[interface{}]bool, 0)
	pre := head
	m[head.Value] = true
	cur := head.Next
	for cur != nil {
		if m[cur.Value] {
			pre.Next = cur.Next
		} else {
			m[cur.Value] = true
			pre = pre.Next
		}
		cur = cur.Next
	}
}

func removeRepetition2(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur := head
	var pre *LNode
	for cur != nil {
		pre = cur //每次对当前cur向后遍历是否有重复
		tmp := cur.Next
		for tmp != nil {
			if tmp.Value == cur.Value {
				pre.Next = tmp.Next
			} else {
				pre = pre.Next
			}
			tmp = tmp.Next
		}
		cur = cur.Next
	}
}
