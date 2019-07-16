package stack

import (
	"btree"
)

//由数组构建最大值树
func getMaxTree(arr []int) *btree.Node {
	l := len(arr)
	var head *btree.Node
	lBigMap := make(map[*btree.Node]*btree.Node, l)
	rBigMap := make(map[*btree.Node]*btree.Node, l)
	st := New() //栈维护一个递增节点集合，便于记录当前节点和前一个比他大的节点
	nodes := make([]*btree.Node, l)
	for k, v := range arr {
		nodes[k] = &btree.Node{Value: v}
	}

	//left 左边第一个比他大的数
	for _, v := range nodes {
		for !st.IsEmpty() && st.Peek().(*btree.Node).Value < v.Value {
			popStackSetMap(st, lBigMap)
		}
		st.Push(v)
	}
	for !st.IsEmpty() {
		popStackSetMap(st, lBigMap)
	}

	//right 右边第一个比他大的数
	for i := l - 1; i >= 0; i-- {
		for !st.IsEmpty() && st.Peek().(*btree.Node).Value < nodes[i].Value {
			popStackSetMap(st, rBigMap)
		}
		st.Push(nodes[i])
	}
	for !st.IsEmpty() {
		popStackSetMap(st, rBigMap)
	}

	for _, v := range nodes {
		left := lBigMap[v]
		right := rBigMap[v]
		if left == nil && right == nil {
			//左右无大值则当前节点为根节点
			head = v
		} else if left == nil {
			if right.Left == nil {
				right.Left = v
			} else {
				right.Right = v
			}
		} else if right == nil {
			if left.Left == nil {
				left.Left = v
			} else {
				left.Right = v
			}
		} else {
			//父节点为当前节点左右大值中较小者
			parent := left
			if left.Value > right.Value {
				parent = right
			}
			if parent.Left == nil {
				parent.Left = v
			} else {
				parent.Right = v
			}
		}
	}
	return head
}

func popStackSetMap(s *S, m map[*btree.Node]*btree.Node) {
	top, _ := s.Pop()
	if s.IsEmpty() {
		m[top.(*btree.Node)] = nil
	} else {
		m[top.(*btree.Node)] = s.Peek().(*btree.Node)
	}
}
