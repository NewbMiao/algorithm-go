package subtree

import (
	"github.com/NewbMiao/algorithm-go/btree/traversal"
	"github.com/NewbMiao/algorithm-go/kit/btree"
	"github.com/NewbMiao/algorithm-go/stringalg"
)

func T1SubtreeEqualT2(t1, t2 *btree.Node) (r bool) {
	st1 := traversal.SerialByPre(t1)
	st2 := traversal.SerialByPre(t2)
	return stringalg.GetIndexOf(st1, st2) > -1
}
