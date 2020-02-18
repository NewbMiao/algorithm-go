package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	l3 := new(ListNode)
	var cur = l3
	var sum, isOut int
	for l1 != nil || l2 != nil {
		sum = isOut
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		isOut = sum / 10
		cur.Next = new(ListNode) //第一节点为空，不反回
		cur = cur.Next
		cur.Val = sum % 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if isOut > 0 {
		cur.Next = new(ListNode)
		cur.Next.Val = 1
	}
	return l3.Next
}
