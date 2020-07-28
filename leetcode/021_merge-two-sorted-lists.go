package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	res := &ListNode{}
	res.Next = &ListNode{}
	cur := res.Next
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				cur.Val = l1.Val
				l1 = l1.Next
			} else {
				cur.Val = l2.Val
				l2 = l2.Next
			}
		}
		if l1 == nil {
			cur.Next = l2
			break
		} else if l2 == nil {
			cur.Next = l1
			break
		}
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	return res.Next
}

func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRecursive(l1, l2.Next)
	return l2
}
