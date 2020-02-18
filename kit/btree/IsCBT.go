package btree

import (
	"github.com/NewbMiao/AlgorithmCodeNote/kit/linkedlist"
)

func IsCBT(bt *Node) bool {
	if bt == nil {
		return true
	}
	var l, r *Node
	leaf := false
	queue := linkedlist.NewList()
	queue.Push(bt)
	for !queue.IsEmpty() {
		tmp := queue.Pop()
		if tmp == nil {
			return false
		}
		node := tmp.(*Node)
		if node == nil {
			return false
		}
		l = node.Left
		r = node.Right
		if (leaf && (l != nil || r != nil) ) || (l == nil && r != nil) {
			return false
		}

		if l != nil {
			queue.Push(l)
		}

		if r != nil {
			queue.Push(r)
		} else {
			leaf = true
		}
	}
	return true
}
