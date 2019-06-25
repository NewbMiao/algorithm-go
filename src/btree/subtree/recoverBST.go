package subtree

import (
	. "btree"
	"stack"
)

func FindTwoErrNode(bt *Node) (errNodes [2]*Node) {
	errNodes = [2]*Node{}
	s := stack.New()
	cur := bt
	var pre *Node
	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			if tmp, ok := s.Pop(); ok && tmp != nil {
				node := tmp.(*Node)
				if node == nil {
					continue
				}
				if pre != nil {
					if pre.Value > node.Value {
						if errNodes[0] == nil {
							errNodes[0] = pre
						}
						errNodes[1] = node
					}
				}
				pre = node
				cur = node.Right

			}
		}
	}
	return
}

func FindTwoErrNodeParent(bt *Node, errNodes [2]*Node) (parentNodes [2]*Node) {
	parentNodes = [2]*Node{}
	s := stack.New()
	cur := bt
	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			if tmp, ok := s.Pop(); ok && tmp != nil {
				node := tmp.(*Node)
				if node == nil {
					continue
				}
				if node.Left == errNodes[0] || node.Right == errNodes[0] {
					parentNodes[0] = node
				}
				if node.Left == errNodes[1] || node.Right == errNodes[1] {
					parentNodes[1] = node
				}
				cur = node.Right
			}
		}
	}
	return
}

func RecoverTree(bt *Node) *Node {
	errs := FindTwoErrNode(bt)
	parents := FindTwoErrNodeParent(bt, errs)
	e1 := errs[0]
	e1P := parents[0]
	e1L := e1.Left
	e1R := e1.Right
	e2 := errs[1]
	e2P := parents[1]
	e2L := e2.Left
	e2R := e2.Right
	//14 situation
	if (e1 == bt) {
		if (e1 == e2P) {
			e1.Left = e2L
			e1.Right = e2R
			e2.Right = e1
			e2.Left = e1L
		} else if (e2P.Left == e2) {
			e2P.Left = e1
			e2.Left = e1L
			e2.Right = e1R
			e1.Left = e2L
			e1.Right = e2R
		} else {
			e2P.Right = e1
			e2.Left = e1L
			e2.Right = e1R
			e1.Left = e2L
			e1.Right = e2R
		}
		bt = e2
	} else if (e2 == bt) {
		if (e2 == e1P) {
			e2.Left = e1L
			e2.Right = e1R
			e1.Left = e2
			e1.Right = e2R
		} else if (e1P.Left == e1) {
			e1P.Left = e2
			e1.Left = e2L
			e1.Right = e2R
			e2.Left = e1L
			e2.Right = e1R
		} else {
			e1P.Right = e2
			e1.Left = e2L
			e1.Right = e2R
			e2.Left = e1L
			e2.Right = e1R
		}
		bt = e1
	} else {
		if (e1 == e2P) {
			if (e1P.Left == e1) {
				e1P.Left = e2
				e1.Left = e2L
				e1.Right = e2R
				e2.Left = e1L
				e2.Right = e1
			} else {
				e1P.Right = e2
				e1.Left = e2L
				e1.Right = e2R
				e2.Left = e1L
				e2.Right = e1
			}
		} else if (e2 == e1P) {
			if (e2P.Left == e2) {
				e2P.Left = e1
				e2.Left = e1L
				e2.Right = e1R
				e1.Left = e2
				e1.Right = e2R
			} else { // ʮ
				e2P.Right = e1
				e2.Left = e1L
				e2.Right = e1R
				e1.Left = e2
				e1.Right = e2R
			}
		} else {
			if (e1P.Left == e1) {
				if (e2P.Left == e2) {
					e1.Left = e2L
					e1.Right = e2R
					e2.Left = e1L
					e2.Right = e1R
					e1P.Left = e2
					e2P.Left = e1
				} else {
					e1.Left = e2L
					e1.Right = e2R
					e2.Left = e1L
					e2.Right = e1R
					e1P.Left = e2
					e2P.Right = e1
				}
			} else {
				if (e2P.Left == e2) {
					e1.Left = e2L
					e1.Right = e2R
					e2.Left = e1L
					e2.Right = e1R
					e1P.Right = e2
					e2P.Left = e1
				} else {
					e1.Left = e2L
					e1.Right = e2R
					e2.Left = e1L
					e2.Right = e1R
					e1P.Right = e2
					e2P.Right = e1
				}
			}
		}
	}
	return bt
}
