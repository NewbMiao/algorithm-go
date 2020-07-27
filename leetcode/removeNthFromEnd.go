package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	nth := head //要删除节点前一个
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
