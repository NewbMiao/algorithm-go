package linkedlist

import (
	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

// 打印有序链表的公共部分.
func printCommonPart(h1, h2 *linkedlist.LList) (res []int) {
	if h1 == nil || h2 == nil {
		return
	}
	n1 := h1.Head
	n2 := h2.Head
	for n1 != nil && n2 != nil {
		if n1.Value.(int) == n2.Value.(int) {
			res = append(res, n1.Value.(int))
			n1 = n1.Next
			n2 = n2.Next
		} else if n1.Value.(int) > n2.Value.(int) {
			n2 = n2.Next
		} else {
			n1 = n1.Next
		}
	}
	return
}
