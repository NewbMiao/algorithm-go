package subtree

import (
	. "kit/btree"
	"testing"
	"fmt"
)

var sbt *Node

func init() {
	sbt = &Node{Value: 6}
	sbt.Left = &Node{Value: 1}
	sbt.Left.Left = &Node{Value: 0}
	sbt.Left.Right = &Node{Value: 3}
	sbt.Right = &Node{Value: 12}
	sbt.Right.Left = &Node{Value: 10}
	sbt.Right.Left.Left = &Node{Value: 4}
	sbt.Right.Left.Left.Left = &Node{Value: 2}
	sbt.Right.Left.Left.Right = &Node{Value: 5}
	sbt.Right.Left.Right = &Node{Value: 14}
	sbt.Right.Left.Right.Left = &Node{Value: 11}
	sbt.Right.Left.Right.Right = &Node{Value: 15}
	sbt.Right.Right = &Node{Value: 13}
	sbt.Right.Right.Left = &Node{Value: 20}
	sbt.Right.Right.Right = &Node{Value: 16}
	sbt.Right.Right.Right.Right = &Node{Value: 23}

	PrintTree(sbt)
}

func TestGetBSTTopSize1(t *testing.T) {
	want := 9
	t.Log(fmt.Sprintf("GetBSTTopSize , want: %v", want))

	r := GetBSTTopSize0(sbt)
	if r == want {
		t.Log("GetBSTTopSize0 is ok")
	} else {
		t.Error("GetBSTTopSize0 is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = GetBSTTopSize1(sbt)
	if r == want {
		t.Log("GetBSTTopSize1 is ok")
	} else {
		t.Error("GetBSTTopSize1 is not ok, result:", fmt.Sprintf("%v", r))
	}

	r = GetBSTTopSize2(sbt)
	if r == want {
		t.Log("GetBSTTopSize2 is ok")
	} else {
		t.Error("GetBSTTopSize2 is not ok, result:", fmt.Sprintf("%v", r))
	}

}
