package subtree

import (
	"fmt"
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var t1, t2 *btree.Node

func init() {
	t1 = &btree.Node{Value: 1}
	t1.Left = &btree.Node{Value: 2}
	t1.Right = &btree.Node{Value: 3}
	t1.Left.Left = &btree.Node{Value: 4}
	t1.Left.Right = &btree.Node{Value: 5}
	t1.Right.Left = &btree.Node{Value: 6}
	t1.Right.Right = &btree.Node{Value: 7}
	t1.Left.Left.Left = &btree.Node{Value: 8}
	t1.Left.Left.Right = &btree.Node{Value: 9}
	t1.Left.Right.Left = &btree.Node{Value: 10}

	t2 = &btree.Node{Value: 2}
	t2.Left = &btree.Node{Value: 4}
	t2.Left.Left = &btree.Node{Value: 8}
	t2.Right = &btree.Node{Value: 5}
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
