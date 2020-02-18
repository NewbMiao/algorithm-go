package linkedlist

import (
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/linkedlist"
)

//判断两链表是否有交点并获取
func getIntersectNode(h1, h2 *LNode) *LNode {
	if h1 == nil || h2 == nil {
		return nil
	}
	l1 := getLoopNode(h1)
	l2 := getLoopNode(h2)

	//有环相交
	if l1 != nil && l2 != nil {
		return bothLoop(h1, h2, l1, l2)
	} else if l1 == nil && l2 == nil { //无环相交
		return noLoop(h1, h2)
	}
	return nil
}

func getLoopNode(head *LNode) *LNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	fast := head.Next.Next
	slow := head.Next
	for fast != slow {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

//无环相交点判定
func noLoop(h1, h2 *LNode) *LNode {
	if h1 == nil || h2 == nil {
		return nil
	}
	n := 0 //计算两无环链表相交前长度差
	cur1 := h1
	cur2 := h2
	for cur1 != nil {
		n++
		cur1 = cur1.Next
	}
	for cur2 != nil {
		n--
		cur2 = cur2.Next
	}
	if n > 0 {
		cur1 = h1
		cur2 = h2
	} else {
		n = -n
		cur1 = h2
		cur2 = h1
	}
	for n != 0 {
		cur1 = cur1.Next
		n--
	}
	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur1
}

//两有环链表是否相交判定
func bothLoop(h1, h2, loop1, loop2 *LNode) *LNode {
	if h1 == nil || h2 == nil || loop1 == nil || loop2 == nil {
		return nil
	}
	if loop1 == loop2 {
		n := 0 //计算两有环链表相交前长度差
		cur1 := h1
		cur2 := h2
		for cur1 != loop1 {
			n++
			cur1 = cur1.Next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.Next
		}
		if n > 0 {
			cur1 = h1
			cur2 = h2
		} else {
			n = -n
			cur1 = h2
			cur2 = h1
		}
		for n != 0 {
			cur1 = cur1.Next
			n--
		}
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		cur := loop1.Next
		for cur != loop1 {
			if cur == loop2 {
				return loop1 //返回 loop1和loop2都可以
			}
			cur = cur.Next
		}
	}

	return nil
}
