package subtree

import (
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/btree"
	"testing"
	"fmt"
)

var et1, et2 *Node

func init() {
	et1 = &Node{Value: 1}
	et1.Left = &Node{Value: 2}
	et1.Right = &Node{Value: 3}
	et1.Left.Left = &Node{Value: 4}
	et1.Left.Right = &Node{Value: 5}
	et1.Right.Left = &Node{Value: 6}
	et1.Right.Right = &Node{Value: 7}
	et1.Left.Left.Right = &Node{Value: 8}
	et1.Left.Right.Left = &Node{Value: 9}

	et2 = &Node{Value: 2}
	et2.Left = &Node{Value: 4}
	et2.Left.Right = &Node{Value: 8}
	et2.Right = &Node{Value: 5}
	et2.Right.Left = &Node{Value: 9}

	PrintTree(et1)
	PrintTree(et2)
}

func TestT1SubtreeEqualT2(t *testing.T) {
	want := true
	t.Log(fmt.Sprintf("T1SubtreeEqualT2 ? want: %v", want))
	r := T1SubtreeEqualT2(et1, et2)
	if r == want {
		t.Log("T1SubtreeEqualT2 is ok")
	} else {
		t.Error("T1SubtreeEqualT2 is not ok, result:", r)
	}
}
