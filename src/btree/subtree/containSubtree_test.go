package subtree

import (
	. "kit/btree"
	"testing"
	"fmt"
)

var t1, t2 *Node

func init() {
	t1 = &Node{Value: 1}
	t1.Left = &Node{Value: 2}
	t1.Right = &Node{Value: 3}
	t1.Left.Left = &Node{Value: 4}
	t1.Left.Right = &Node{Value: 5}
	t1.Right.Left = &Node{Value: 6}
	t1.Right.Right = &Node{Value: 7}
	t1.Left.Left.Left = &Node{Value: 8}
	t1.Left.Left.Right = &Node{Value: 9}
	t1.Left.Right.Left = &Node{Value: 10}

	t2 = &Node{Value: 2}
	t2.Left = &Node{Value: 4}
	t2.Left.Left = &Node{Value: 8}
	t2.Right = &Node{Value: 5}
}

func TestT1ContainsT2(t *testing.T) {
	want := true
	t.Log(fmt.Sprintf("T1ContainsT2 ? want: %v", want))
	r := T1ContainsT2(t1, t2)
	if r == want {
		t.Log("T1ContainsT2 is ok")
	} else {
		t.Error("T1ContainsT2 is not ok, result:", r)
	}
}
