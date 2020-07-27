package subtree

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var sbt *btree.Node

func init() {
	sbt = &btree.Node{Value: 6}
	sbt.Left = &btree.Node{Value: 1}
	sbt.Left.Left = &btree.Node{Value: 0}
	sbt.Left.Right = &btree.Node{Value: 3}
	sbt.Right = &btree.Node{Value: 12}
	sbt.Right.Left = &btree.Node{Value: 10}
	sbt.Right.Left.Left = &btree.Node{Value: 4}
	sbt.Right.Left.Left.Left = &btree.Node{Value: 2}
	sbt.Right.Left.Left.Right = &btree.Node{Value: 5}
	sbt.Right.Left.Right = &btree.Node{Value: 14}
	sbt.Right.Left.Right.Left = &btree.Node{Value: 11}
	sbt.Right.Left.Right.Right = &btree.Node{Value: 15}
	sbt.Right.Right = &btree.Node{Value: 13}
	sbt.Right.Right.Left = &btree.Node{Value: 20}
	sbt.Right.Right.Right = &btree.Node{Value: 16}
	sbt.Right.Right.Right.Right = &btree.Node{Value: 23}

	btree.PrintTree(sbt)
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
