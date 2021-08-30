package linkedlist

import (
	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
	"github.com/NewbMiao/algorithm-go/kit/stack"
)

// 每k个节点逆置.
func reverseKNodes1(head *linkedlist.LNode, k int) *linkedlist.LNode {
	if k < 2 {
		return head
	}
	if head == nil {
		return nil
	}
	var newHead, newCur, next *linkedlist.LNode
	st := stack.New()
	cur := head
	newHead = head
	n := 0
	for cur != nil {
		st.Push(cur)
		next = cur.Next
		n++
		if n == k {
			for !st.IsEmpty() {
				top, _ := st.Pop()
				node, _ := top.(*linkedlist.LNode)
				if newHead == head {
					newHead = node
					newCur = node
				} else {
					newCur.Next = node
					newCur = newCur.Next
				}
				n--
			}
			newCur.Next = next
		}
		cur = next
	}
	return newHead
}

func reverseKNodes2(head *linkedlist.LNode, k int) *linkedlist.LNode {
	if k < 2 {
		return head
	}
	if head == nil {
		return nil
	}
	cur := head
	newHead := head
	var start *linkedlist.LNode
	var left, next *linkedlist.LNode
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

func resignList(left, start, end, right *linkedlist.LNode) {
	pre := start
	cur := start.Next
	var next *linkedlist.LNode
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

func reverseKNodes3(head *linkedlist.LNode, k int) *linkedlist.LNode {
	if k < 2 {
		return head
	}
	if head == nil {
		return nil
	}
	tmp := head
	for i := 1; i < k && tmp != nil; i++ {
		tmp = tmp.Next
	}
	if tmp == nil {
		return head
	}

	// 下一组开始节点
	tmp2 := tmp.Next
	tmp.Next = nil // 移除next
	// 逆置当前分组
	newHead := reverseList(head)

	// 递归后续分组
	nextGroup := reverseKNodes3(tmp2, k)
	head.Next = nextGroup

	return newHead
}

// func reverseList(head *linkedlist.LNode) *linkedlist.LNode {
//	var pre, next *linkedlist.LNode
//	for head != nil {
//		next = head.Next
//		head.Next = pre
//		pre = head
//		head = next
//	}
//	return pre
// }

func reverseList(head *linkedlist.LNode) *linkedlist.LNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}
