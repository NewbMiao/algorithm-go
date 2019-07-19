package subtree

import (
	. "kit/btree"
	"btree/traversal"
	"string_alg"
)

func T1SubtreeEqualT2(t1, t2 *Node) (r bool) {
	st1 := traversal.SerialByPre(t1)
	st2 := traversal.SerialByPre(t2)
	return string_alg.GetIndexOf(st1, st2) > -1
}
