package linkedlist

import (
	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
	"github.com/NewbMiao/algorithm-go/kit/stack"
)

func removeValue1(head *linkedlist.LNode, num int) *linkedlist.LNode {
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
		st.Peek().(*linkedlist.LNode).Next = head
		top, _ := st.Pop()
		head = top.(*linkedlist.LNode)
	}
	return head
}

func removeValue2(head *linkedlist.LNode, num int) *linkedlist.LNode {
	if head == nil {
		return nil
	}

	var newHead, pre *linkedlist.LNode
	cur := head

	for cur != nil {
		if cur.Value.(int) != num {
			if newHead == nil {
				newHead = cur
			}
			pre = cur
		} else if pre != nil {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return newHead
}
