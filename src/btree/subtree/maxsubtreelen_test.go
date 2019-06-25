package subtree

import (
	"testing"
	"fmt"
	. "btree"
)

var sbt *Node

func init() {
	sbt = &Node{Value: -3}
	sbt.Left = &Node{Value: 3}
	sbt.Right = &Node{Value: -9}
	sbt.Left.Left = &Node{Value: 1}
	sbt.Left.Right = &Node{Value: 0}
	sbt.Left.Right.Left = &Node{Value: 1}
	sbt.Left.Right.Right = &Node{Value: 6}
	sbt.Right.Left = &Node{Value: 2}
	sbt.Right.Right = &Node{Value: 1}
	PrintTree(sbt)
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
		r := GetMaxSubtreeLen(sbt, sum)
		return
		if r == want {
			t.Log("GetMaxSubtreeLen is ok")
		} else {
			t.Error("GetMaxSubtreeLen is not ok, result:", fmt.Sprintf("%v", r))
		}
	}

}
