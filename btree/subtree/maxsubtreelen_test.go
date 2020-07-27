package subtree

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var smbt *btree.Node

func init() {
	smbt = &btree.Node{Value: -3}
	smbt.Left = &btree.Node{Value: 3}
	smbt.Right = &btree.Node{Value: -9}
	smbt.Left.Left = &btree.Node{Value: 1}
	smbt.Left.Right = &btree.Node{Value: 0}
	smbt.Left.Right.Left = &btree.Node{Value: 1}
	smbt.Left.Right.Right = &btree.Node{Value: 6}
	smbt.Right.Left = &btree.Node{Value: 2}
	smbt.Right.Right = &btree.Node{Value: 1}
	btree.PrintTree(smbt)
}

func TestGetMaxSubtreeLen(t *testing.T) {
	wants := map[int]int{
		6:   4,
		-9:  1,
		-12: 2,
		1:   4,
		2:   1,
	}
	for sum, want := range wants {
		t.Log(fmt.Sprintf("GetMaxSubtreeLen sum:%d, want: %v", sum, want))
		r := GetMaxSubtreeLen(smbt, sum)
		if r == want {
			t.Log("GetMaxSubtreeLen is ok")
		} else {
			t.Error("GetMaxSubtreeLen is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
