package btree

import (
	"fmt"
	"testing"
)

var cbt *Node

func init() {
	cbt = &Node{Value: 4}
	cbt.Left = &Node{Value: 2}
	cbt.Right = &Node{Value: 6}
	cbt.Left.Left = &Node{Value: 1}
	cbt.Left.Right = &Node{Value: 3}
	cbt.Right.Left = &Node{Value: 5}
}

func TestIsCBT(t *testing.T) {
	want := true
	t.Log(fmt.Sprintf("IsCBT ? want: %v", want))
	r := IsCBT(cbt)
	if r == want {
		t.Log("IsCBT is ok")
	} else {
		t.Error("IsCBT is not ok, result:", r)
	}
	PrintTree(cbt)
}
