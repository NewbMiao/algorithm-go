package traversal

import (
	"testing"

	"github.com/NewbMiao/algorithm-go/kit/btree"
)

var pbt *btree.Node

func init() {
	pbt = &btree.Node{Value: 1}
	pbt.Left = &btree.Node{Value: 2}
	pbt.Right = &btree.Node{Value: 3}
	pbt.Left.Left = &btree.Node{Value: 4}
	pbt.Right.Left = &btree.Node{Value: 5}
	pbt.Right.Right = &btree.Node{Value: 6}
	pbt.Right.Left.Left = &btree.Node{Value: 7}
	pbt.Right.Left.Right = &btree.Node{Value: 8}

	btree.PrintTree(pbt)
}

func TestPrintByLevel(t *testing.T) {
	PrintByLevel(pbt)
}

func TestPrintByZigZag(t *testing.T) {
	PrintByZigZag(pbt)
}
