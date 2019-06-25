package traversal

import (
	"stack"
	. "btree"
)

//先序遍历：根,左,右
func PreOrderRecur(bt *Node) (res []int) {
	if bt == nil {
		return
	}
	res = append(res, bt.Value)
	res = append(res, PreOrderRecur(bt.Left)...)
	res = append(res, PreOrderRecur(bt.Right)...)
	return
}

//中序遍历：左,根,右
func InOrderRecur(bt *Node) (res []int) {
	if bt == nil {
		return
	}
	res = append(res, InOrderRecur(bt.Left)...)
	res = append(res, bt.Value)
	res = append(res, InOrderRecur(bt.Right)...)
	return
}

//后序遍历：左,右,根
func LastOrderRecur(bt *Node) (res []int) {
	if bt == nil {
		return
	}
	res = append(res, LastOrderRecur(bt.Left)...)
	res = append(res, LastOrderRecur(bt.Right)...)
	res = append(res, bt.Value)
	return
}

func PreOrder(bt *Node) (res []int) {
	if bt == nil {
		return
	}
	s := &stack.S{}
	s.Push(bt)
	for !s.IsEmpty() {
		tmp, _ := s.Pop()
		cur := tmp.(*Node)
		res = append(res, cur.Value)
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
	return
}

func InOrder(bt *Node) (res []int) {
	if bt == nil {
		return
	}
	s := &stack.S{}
	cur := bt
	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			//添加二叉树左边界
			s.Push(cur)
			cur = cur.Left
		} else {
			//左节点为空，切到当前节点的右节点
			tmp, ok := s.Pop()
			if !ok {
				return
			}
			node := tmp.(*Node)
			res = append(res, node.Value)
			cur = node.Right
		}
	}
	return
}
func LastOrder1(bt *Node) (res []int) {
	if bt == nil {
		return
	}
	s1 := &stack.S{} //记录 根 左右 入栈， 根 右左 出栈顺序
	s2 := &stack.S{} //逆序s1, 根右左 入栈,记录 左右根 出栈顺序
	s1.Push(bt)
	for !s1.IsEmpty() {
		if tmp, ok := s1.Pop(); ok {
			node := tmp.(*Node)
			s2.Push(node.Value)
			if node.Left != nil {
				s1.Push(node.Left)
			}
			if node.Right != nil {
				s1.Push(node.Right)
			}
		}

	}
	for !s2.IsEmpty() {
		tmp, _ := s2.Pop()
		res = append(res, tmp.(int))
	}
	return
}

func LastOrder2(bt *Node) (res []int) {
	if bt == nil {
		return
	}
	var history *Node //最近一次弹出并打印节点
	var cur *Node     //当前栈顶节点
	s := &stack.S{}
	s.Push(bt)
	history = bt //history为nil会影响最左节点的添加
	for !s.IsEmpty() {
		cur = s.Peek().(*Node)
		if cur == nil {
			return
		}

		if cur.Left != nil && cur.Left != history && cur.Right != history {
			//左子树未处理
			s.Push(cur.Left)
		} else if cur.Right != nil && cur.Right != history {
			//右子树未处理
			s.Push(cur.Right)
		} else {
			//都处理过
			s.Pop()
			history = cur
			res = append(res, history.Value)
		}

	}
	return
}
