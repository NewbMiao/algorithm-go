package leetcode



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
