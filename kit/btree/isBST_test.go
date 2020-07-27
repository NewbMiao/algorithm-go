package btree

import (
	"fmt"
	"testing"
)

var bst *Node

func init() {
	bst = &Node{Value: 4}
	bst.Left = &Node{Value: 2}
	bst.Right = &Node{Value: 6}
	bst.Left.Left = &Node{Value: 1}
	bst.Left.Right = &Node{Value: 3}
	bst.Right.Left = &Node{Value: 5}
}
func TestIsBST(t *testing.T) {
	want := true
	t.Log(fmt.Sprintf("IsBST ? want: %v", want))
	r := IsBST(bst)
	if r == want {
		t.Log("IsBST is ok")
	} else {
		t.Error("IsBST is not ok, result:", r)
	}
	PrintTree(bst)
}
