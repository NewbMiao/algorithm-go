package linkedlist

import (
	. "kit/linkedlist"
	"kit/stack"
)

//每k个节点逆置
func reverseKNodes1(head *LNode, k int) *LNode {
	if k < 2 {
		return head
	}
	if head == nil {
		return nil
	}
	var newHead, newCur, nextGroup *LNode
	st := stack.New()
	cur := head
	n := 0
	for cur != nil {
		if n == 0 {
			nextGroup = cur
		}
		st.Push(cur)
		n++
		cur = cur.Next //先移动cur，避免逆置后获取错误的cur
		if n == k {
			for !st.IsEmpty() {
				top, _ := st.Pop()
				node, _ := top.(*LNode)
				if newHead == nil {
					newHead = node
					newCur = node
				} else {
					newCur.Next = node
					newCur = newCur.Next
				}
				n--
			}
			newCur.Next = nil
		}

	}
	if n > 0 {
		if newCur != nil {
			newCur.Next = nextGroup
		} else {
			return head
		}
	}
	return newHead
}

func reverseKNodes2(head *LNode, k int) *LNode {
	if k < 2 {
		return head
	}
	if head == nil {
		return nil
	}
	cur := head
	newHead := head
	start := head
	var left, next *LNode
	n := 1
	for cur != nil {
		next = cur.Next
		if n == k {
			if left == nil {
				start = head
			} else {
				start = left.Next
			}
			resignList(left, start, cur, next)
			if newHead == head {
				newHead = cur
			}
			left = start
			n = 0
		}
		n++
		cur = next
	}
	return newHead
}

func resignList(left, start, end, right *LNode) {
	pre := start
	cur := start.Next
	var next *LNode
	for cur != right {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	if left != nil {
		left.Next = end
	}
	start.Next = right
}
