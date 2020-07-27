package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curSum, l3 := 0, new(ListNode)
	for curNode := l3; l1 != nil || l2 != nil || curSum > 0; curNode = curNode.Next {
		if l1 != nil {
			curSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curSum += l2.Val
			l2 = l2.Next
		}
		curNode.Next = &ListNode{curSum % 10, nil}
		curSum /= 10
	}
	return l3.Next
}
