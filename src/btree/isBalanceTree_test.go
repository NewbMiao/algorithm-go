package btree

import (
	"testing"
	"fmt"
)

var bbt *Node

func init() {
	bbt = &Node{Value: 1}
	bbt.Left = &Node{Value: 2}
	bbt.Right = &Node{Value: 3}
	bbt.Left.Left = &Node{Value: 4}
	bbt.Left.Right = &Node{Value: 5}
	bbt.Right.Left = &Node{Value: 6}
	bbt.Right.Right = &Node{Value: 7}

}
func TestIsBalanceTree(t *testing.T) {
	want := true
	t.Log(fmt.Sprintf("IsBalanceTree ? want: %v", want))
	r := IsBalanceTree(bbt)
	if r == want {
		t.Log("IsBalanceTree is ok")
	} else {
		t.Error("IsBalanceTree is not ok, result:", r)
	}
	PrintTree(bbt)
}
