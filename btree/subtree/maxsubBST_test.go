package subtree

import (
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/btree"
	"testing"
)

var mbt1 *Node

func init() {
	mbt1 = &Node{Value: 6}
	mbt1.Left = &Node{Value: 1}
	mbt1.Left.Left = &Node{Value: 0}
	mbt1.Left.Right = &Node{Value: 3}
	mbt1.Right = &Node{Value: 12}
	mbt1.Right.Left = &Node{Value: 10}
	mbt1.Right.Left.Left = &Node{Value: 4}
	mbt1.Right.Left.Left.Left = &Node{Value: 2}
	mbt1.Right.Left.Left.Right = &Node{Value: 5}
	mbt1.Right.Left.Right = &Node{Value: 14}
	mbt1.Right.Left.Right.Left = &Node{Value: 11}
	mbt1.Right.Left.Right.Right = &Node{Value: 15}
	mbt1.Right.Left.Right.Right.Right = &Node{Value: 19}
	mbt1.Right.Right = &Node{Value: 13}
	mbt1.Right.Right.Left = &Node{Value: 20}
	mbt1.Right.Right.Right = &Node{Value: 16}
	PrintTree(mbt1)
}

func TestGetMaxSubBST(t *testing.T) {
	r := GetMaxSubBST(mbt1)
	PrintTree(r)

}
