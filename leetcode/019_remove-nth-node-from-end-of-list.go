package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	nth := head // 要删除节点前一个
	llen := 0
	for cur != nil {
		cur = cur.Next
		if llen > n {
			nth = nth.Next
		}
		llen++
	}
	if llen > n {
		nth.Next = nth.Next.Next
	} else if llen == n {
		nth = nth.Next
		return nth
	} else {
		return nil
	}
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	first, last := dummyHead, dummyHead
	for i := 0; i < n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		last = last.Next
	}
	last.Next = last.Next.Next
	return dummyHead.Next
}
