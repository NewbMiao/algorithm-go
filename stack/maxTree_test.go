package stack

import (
	"testing"
	"github.com/NewbMiao/AlgorithmCodeNote/btree/traversal"
	"fmt"
	"github.com/NewbMiao/AlgorithmCodeNote/kit/btree"
)

func TestGetMaxTree(t *testing.T) {
	d := []int{3, 4, 5, 1, 2}
	bt := getMaxTree(d)
	res := [][]int{}
	res = append(res, traversal.PreOrder(bt))
	res = append(res, traversal.InOrder(bt))
	wants := [][]int{
		{5, 4, 3, 2, 1},
		{3, 4, 5, 1, 2},
	}
	t.Log(fmt.Sprintf("getMaxTree, want preOrder and inOrder: %v", wants))

	if fmt.Sprintf("%v", res) == fmt.Sprintf("%v", wants) {
		t.Log("getMaxTree is ok")
	} else {
		t.Error("getMaxTree is not ok, result:", fmt.Sprintf("%v", res))
	}

	btree.PrintTree(bt)

}
