package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode-cn.com/problems/merge-two-sorted-lists/
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。.
*/
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
