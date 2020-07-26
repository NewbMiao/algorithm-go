package subtree

import (
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var mbt1 *btree.Node

func init() {
	mbt1 = &btree.Node{Value: 6}
	mbt1.Left = &btree.Node{Value: 1}
	mbt1.Left.Left = &btree.Node{Value: 0}
	mbt1.Left.Right = &btree.Node{Value: 3}
	mbt1.Right = &btree.Node{Value: 12}
	mbt1.Right.Left = &btree.Node{Value: 10}
	mbt1.Right.Left.Left = &btree.Node{Value: 4}
	mbt1.Right.Left.Left.Left = &btree.Node{Value: 2}
	mbt1.Right.Left.Left.Right = &btree.Node{Value: 5}
	mbt1.Right.Left.Right = &btree.Node{Value: 14}
	mbt1.Right.Left.Right.Left = &btree.Node{Value: 11}
	mbt1.Right.Left.Right.Right = &btree.Node{Value: 15}
	mbt1.Right.Left.Right.Right.Right = &btree.Node{Value: 19}
	mbt1.Right.Right = &btree.Node{Value: 13}
	mbt1.Right.Right.Left = &btree.Node{Value: 20}
	mbt1.Right.Right.Right = &btree.Node{Value: 16}
	btree.PrintTree(mbt1)
}

func TestGetMaxSubBST(t *testing.T) {
	r := GetMaxSubBST(mbt1)
	btree.PrintTree(r)
}
