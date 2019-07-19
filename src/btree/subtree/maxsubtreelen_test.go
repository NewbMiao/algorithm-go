package subtree

import (
	"testing"
	"fmt"
	. "kit/btree"
)

var smbt *Node

func init() {
	smbt = &Node{Value: -3}
	smbt.Left = &Node{Value: 3}
	smbt.Right = &Node{Value: -9}
	smbt.Left.Left = &Node{Value: 1}
	smbt.Left.Right = &Node{Value: 0}
	smbt.Left.Right.Left = &Node{Value: 1}
	smbt.Left.Right.Right = &Node{Value: 6}
	smbt.Right.Left = &Node{Value: 2}
	smbt.Right.Right = &Node{Value: 1}
	PrintTree(smbt)
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
		return
		if r == want {
			t.Log("GetMaxSubtreeLen is ok")
		} else {
			t.Error("GetMaxSubtreeLen is not ok, result:", fmt.Sprintf("%v", r))
		}
	}

}
