package subtree

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var et1, et2 *btree.Node

func init() {
	et1 = &btree.Node{Value: 1}
	et1.Left = &btree.Node{Value: 2}
	et1.Right = &btree.Node{Value: 3}
	et1.Left.Left = &btree.Node{Value: 4}
	et1.Left.Right = &btree.Node{Value: 5}
	et1.Right.Left = &btree.Node{Value: 6}
	et1.Right.Right = &btree.Node{Value: 7}
	et1.Left.Left.Right = &btree.Node{Value: 8}
	et1.Left.Right.Left = &btree.Node{Value: 9}

	et2 = &btree.Node{Value: 2}
	et2.Left = &btree.Node{Value: 4}
	et2.Left.Right = &btree.Node{Value: 8}
	et2.Right = &btree.Node{Value: 5}
	et2.Right.Left = &btree.Node{Value: 9}

	btree.PrintTree(et1)
	btree.PrintTree(et2)
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
