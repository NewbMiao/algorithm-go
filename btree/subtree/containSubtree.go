package subtree

import (
	"github.com/NewbMiao/algorithm-go/kit/btree"
)

//前序遍历
func T1ContainsT2(t1, t2 *btree.Node) bool {
	return check(t1, t2) || T1ContainsT2(t1.Left, t2) || T1ContainsT2(t1.Right, t2)
}

func check(h, t2 *btree.Node) bool {
	if t2 == nil {
		return true
	}
	if h == nil || h.Value != t2.Value {
		return false
	}
	return check(h.Left, t2.Left) && check(h.Right, t2.Right)
}
