package linkedlist

import (
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/linkedlist"
	"github.com/NewbMiao/AlgorithmCodeNote/kit/stack"
)

func removeValue1(head *LNode, num int) *LNode {
	if head == nil {
		return nil
	}
	st := stack.New()
	for head != nil {
		if head.Value.(int) != num {
			st.Push(head)
		}
		head = head.Next
	}

	//倒着重新连接链表
	for !st.IsEmpty() {
		st.Peek().(*LNode).Next = head
		top, _ := st.Pop()
		head = top.(*LNode)
	}
	return head
}

func removeValue2(head *LNode, num int) *LNode {

	if head == nil {
		return nil
	}

	var newHead, pre *LNode
	cur := head

	for cur != nil {
		if cur.Value.(int) != num {
			if newHead == nil {
				newHead = cur
			}
			pre = cur
		} else {
			if pre != nil {
				pre.Next = cur.Next
			}
		}
		cur = cur.Next
	}
	return newHead
}
