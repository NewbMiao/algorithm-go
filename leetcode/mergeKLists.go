package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode-cn.com/problems/merge-k-sorted-lists/
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
*/
// 分治.
func mergeKLists(lists []*ListNode) *ListNode {
	amout := len(lists)
	if amout == 0 {
		return nil
	}
	interval := 1
	for interval < amout {
		//两两合并
		for i := 0; i+interval < amout; i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[interval+i])
			fmt.Println(i, interval+i)
		}
		interval *= 2
		fmt.Println(interval)
	}

	return lists[0]
}
