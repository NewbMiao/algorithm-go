package subtree

import (
	. "kit/btree"
	"fmt"
)

//前序遍历
func T1ContainsT2(t1, t2 *Node) bool {
	return check(t1, t2) || T1ContainsT2(t1.Left, t2) || T1ContainsT2(t1.Right, t2)
}

func check(h, t2 *Node) bool {
	if t2 == nil {
		return true
	}
	fmt.Println(h.Value, t2.Value)

	if h == nil || h.Value != t2.Value {
		return false
	}
	return check(h.Left, t2.Left) && check(h.Right, t2.Right)
}
